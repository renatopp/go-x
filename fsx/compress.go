package fsx

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ErrUnknownFormat is returned by Decompress when the source file does not
// match any of the supported archive or compression formats.
var ErrUnknownFormat = errors.New("unknown archive or compression format")

// ----------------------------------------------------------------------------
// ZIP
// ----------------------------------------------------------------------------

// Zip compresses the source file or directory into a zip archive at the
// destination path. If src is a directory, its contents are added to the
// archive using paths relative to src; the directory itself is not nested
// inside the archive.
func Zip(src, dest string) error {
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	return filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name, err := relArchiveName(src, path, fi)
		if err != nil {
			return err
		}
		if name == "" {
			return nil
		}

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		header.Name = name

		if fi.IsDir() {
			header.Name += "/"
			_, err := zw.CreateHeader(header)
			return err
		}

		header.Method = zip.Deflate
		w, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(w, file)
		return err
	})
}

// Unzip extracts the contents of a zip archive at the source path into the
// destination directory.
func Unzip(src, dest string) error {
	r, err := OpenZip(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err := extractZipEntry(f, dest); err != nil {
			return err
		}
	}
	return nil
}

// OpenZip opens the zip archive at the specified path for reading.
func OpenZip(path string) (*zip.ReadCloser, error) {
	return zip.OpenReader(path)
}

// ZipWriter writes a zip archive to a file, closing both the zip stream and
// the underlying file on Close.
type ZipWriter struct {
	*zip.Writer
	f *os.File
}

// CreateZip creates a new zip archive at the specified path for writing.
func CreateZip(path string) (*ZipWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &ZipWriter{zip.NewWriter(f), f}, nil
}

// Close flushes and closes the zip stream, then closes the underlying file.
func (w *ZipWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// TAR
// ----------------------------------------------------------------------------

// Tar archives the source file or directory into a tar archive at the
// destination path. If src is a directory, its contents are added to the
// archive using paths relative to src; the directory itself is not nested
// inside the archive.
func Tar(src, dest string) error {
	tw, err := CreateTar(dest)
	if err != nil {
		return err
	}
	defer tw.Close()

	return writeTarEntries(tw.Writer, src)
}

// Untar extracts the contents of a tar archive at the source path into the
// destination directory.
func Untar(src, dest string) error {
	tr, err := OpenTar(src)
	if err != nil {
		return err
	}
	defer tr.Close()

	return extractTarEntries(tr.Reader, dest)
}

// TarReader reads a tar archive from a file, closing both the tar stream and
// the underlying file on Close.
type TarReader struct {
	*tar.Reader
	f *os.File
}

// OpenTar opens the tar archive at the specified path for reading.
func OpenTar(path string) (*TarReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &TarReader{tar.NewReader(f), f}, nil
}

// Close closes the underlying file.
func (r *TarReader) Close() error {
	return r.f.Close()
}

// TarWriter writes a tar archive to a file, closing both the tar stream and
// the underlying file on Close.
type TarWriter struct {
	*tar.Writer
	f *os.File
}

// CreateTar creates a new tar archive at the specified path for writing.
func CreateTar(path string) (*TarWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &TarWriter{tar.NewWriter(f), f}, nil
}

// Close flushes and closes the tar stream, then closes the underlying file.
func (w *TarWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// TAR+GZIP
// ----------------------------------------------------------------------------

// TarGz archives the source file or directory into a gzip-compressed tar
// archive at the destination path, combining Tar and Gzip in a single pass.
// If src is a directory, its contents are added to the archive using paths
// relative to src; the directory itself is not nested inside the archive.
func TarGz(src, dest string) error {
	tw, err := CreateTarGz(dest)
	if err != nil {
		return err
	}
	defer tw.Close()

	return writeTarEntries(tw.Writer, src)
}

// UntarGz extracts the contents of a gzip-compressed tar archive at the
// source path into the destination directory.
func UntarGz(src, dest string) error {
	tr, err := OpenTarGz(src)
	if err != nil {
		return err
	}
	defer tr.Close()

	return extractTarEntries(tr.Reader, dest)
}

// TarGzReader reads a gzip-compressed tar archive from a file, closing the
// tar stream, the gzip stream and the underlying file on Close.
type TarGzReader struct {
	*tar.Reader
	gz *gzip.Reader
	f  *os.File
}

// OpenTarGz opens the gzip-compressed tar archive at the specified path for
// reading.
func OpenTarGz(path string) (*TarGzReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &TarGzReader{tar.NewReader(gz), gz, f}, nil
}

// Close closes the gzip stream and the underlying file.
func (r *TarGzReader) Close() error {
	if err := r.gz.Close(); err != nil {
		r.f.Close()
		return err
	}
	return r.f.Close()
}

// TarGzWriter writes a gzip-compressed tar archive to a file, closing the tar
// stream, the gzip stream and the underlying file on Close.
type TarGzWriter struct {
	*tar.Writer
	gz *gzip.Writer
	f  *os.File
}

// CreateTarGz creates a new gzip-compressed tar archive at the specified path
// for writing.
func CreateTarGz(path string) (*TarGzWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	gz := gzip.NewWriter(f)
	return &TarGzWriter{tar.NewWriter(gz), gz, f}, nil
}

// Close flushes and closes the tar stream and the gzip stream, then closes
// the underlying file.
func (w *TarGzWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.gz.Close()
		w.f.Close()
		return err
	}
	if err := w.gz.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// GZIP
// ----------------------------------------------------------------------------

// Gzip compresses the source file into a gzip-compressed file at the
// destination path. src must be a file; if it is a directory, it returns
// ErrIsDir.
func Gzip(src, dest string) error {
	if IsDir(src) {
		return ErrIsDir
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	gw, err := CreateGzip(dest)
	if err != nil {
		return err
	}
	defer gw.Close()

	_, err = io.Copy(gw, in)
	return err
}

// Ungzip decompresses the gzip-compressed source file into the destination
// file.
func Ungzip(src, dest string) error {
	gr, err := OpenGzip(src)
	if err != nil {
		return err
	}
	defer gr.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, gr)
	return err
}

// GzipReader reads gzip-compressed data from a file, closing both the gzip
// stream and the underlying file on Close.
type GzipReader struct {
	*gzip.Reader
	f *os.File
}

// OpenGzip opens the gzip-compressed file at the specified path for reading.
func OpenGzip(path string) (*GzipReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &GzipReader{gz, f}, nil
}

// Close closes the gzip stream and the underlying file.
func (r *GzipReader) Close() error {
	if err := r.Reader.Close(); err != nil {
		r.f.Close()
		return err
	}
	return r.f.Close()
}

// GzipWriter writes gzip-compressed data to a file, closing both the gzip
// stream and the underlying file on Close.
type GzipWriter struct {
	*gzip.Writer
	f *os.File
}

// CreateGzip creates a new gzip-compressed file at the specified path for
// writing.
func CreateGzip(path string) (*GzipWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &GzipWriter{gzip.NewWriter(f), f}, nil
}

// Close flushes and closes the gzip stream, then closes the underlying file.
func (w *GzipWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// ZLIB
// ----------------------------------------------------------------------------

// Zlib compresses the source file into a zlib-compressed file at the
// destination path. src must be a file; if it is a directory, it returns
// ErrIsDir.
func Zlib(src, dest string) error {
	if IsDir(src) {
		return ErrIsDir
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	zw, err := CreateZlib(dest)
	if err != nil {
		return err
	}
	defer zw.Close()

	_, err = io.Copy(zw, in)
	return err
}

// Unzlib decompresses the zlib-compressed source file into the destination
// file.
func Unzlib(src, dest string) error {
	zr, err := OpenZlib(src)
	if err != nil {
		return err
	}
	defer zr.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, zr)
	return err
}

// ZlibReader reads zlib-compressed data from a file, closing both the zlib
// stream and the underlying file on Close.
type ZlibReader struct {
	io.ReadCloser
	f *os.File
}

// OpenZlib opens the zlib-compressed file at the specified path for reading.
func OpenZlib(path string) (*ZlibReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	zr, err := zlib.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &ZlibReader{zr, f}, nil
}

// Close closes the zlib stream and the underlying file.
func (r *ZlibReader) Close() error {
	if err := r.ReadCloser.Close(); err != nil {
		r.f.Close()
		return err
	}
	return r.f.Close()
}

// ZlibWriter writes zlib-compressed data to a file, closing both the zlib
// stream and the underlying file on Close.
type ZlibWriter struct {
	*zlib.Writer
	f *os.File
}

// CreateZlib creates a new zlib-compressed file at the specified path for
// writing.
func CreateZlib(path string) (*ZlibWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &ZlibWriter{zlib.NewWriter(f), f}, nil
}

// Close flushes and closes the zlib stream, then closes the underlying file.
func (w *ZlibWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// FLATE
// ----------------------------------------------------------------------------

// Flate compresses the source file into a raw DEFLATE-compressed file at the
// destination path. src must be a file; if it is a directory, it returns
// ErrIsDir.
func Flate(src, dest string) error {
	if IsDir(src) {
		return ErrIsDir
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	fw, err := CreateFlate(dest)
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, in)
	return err
}

// Unflate decompresses the raw DEFLATE-compressed source file into the
// destination file.
func Unflate(src, dest string) error {
	fr, err := OpenFlate(src)
	if err != nil {
		return err
	}
	defer fr.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, fr)
	return err
}

// FlateReader reads raw DEFLATE-compressed data from a file, closing both the
// flate stream and the underlying file on Close.
type FlateReader struct {
	io.ReadCloser
	f *os.File
}

// OpenFlate opens the raw DEFLATE-compressed file at the specified path for
// reading.
func OpenFlate(path string) (*FlateReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FlateReader{flate.NewReader(f), f}, nil
}

// Close closes the flate stream and the underlying file.
func (r *FlateReader) Close() error {
	if err := r.ReadCloser.Close(); err != nil {
		r.f.Close()
		return err
	}
	return r.f.Close()
}

// FlateWriter writes raw DEFLATE-compressed data to a file, closing both the
// flate stream and the underlying file on Close.
type FlateWriter struct {
	*flate.Writer
	f *os.File
}

// CreateFlate creates a new raw DEFLATE-compressed file at the specified path
// for writing.
func CreateFlate(path string) (*FlateWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	fw, err := flate.NewWriter(f, flate.DefaultCompression)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &FlateWriter{fw, f}, nil
}

// Close flushes and closes the flate stream, then closes the underlying file.
func (w *FlateWriter) Close() error {
	if err := w.Writer.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// LZW
// ----------------------------------------------------------------------------

// DefaultLzwOrder and DefaultLzwLitWidth are the bit ordering and literal
// code width used by Lzw, Unlzw, OpenLzw and CreateLzw.
const (
	DefaultLzwOrder    = lzw.MSB
	DefaultLzwLitWidth = 8
)

// Lzw compresses the source file into an LZW-compressed file at the
// destination path, using DefaultLzwOrder and DefaultLzwLitWidth. src must be
// a file; if it is a directory, it returns ErrIsDir.
func Lzw(src, dest string) error {
	return LzwWith(src, dest, DefaultLzwOrder, DefaultLzwLitWidth)
}

// LzwWith is like Lzw but allows the bit ordering and literal code width to
// be specified explicitly. litWidth must be in the range [2,8] and must match
// the litWidth used to decompress the file.
func LzwWith(src, dest string, order lzw.Order, litWidth int) error {
	if IsDir(src) {
		return ErrIsDir
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	lw, err := CreateLzwWith(dest, order, litWidth)
	if err != nil {
		return err
	}
	defer lw.Close()

	_, err = io.Copy(lw, in)
	return err
}

// Unlzw decompresses the LZW-compressed source file into the destination
// file, using DefaultLzwOrder and DefaultLzwLitWidth.
func Unlzw(src, dest string) error {
	return UnlzwWith(src, dest, DefaultLzwOrder, DefaultLzwLitWidth)
}

// UnlzwWith is like Unlzw but allows the bit ordering and literal code width
// to be specified explicitly. They must match the values used to compress the
// file.
func UnlzwWith(src, dest string, order lzw.Order, litWidth int) error {
	lr, err := OpenLzwWith(src, order, litWidth)
	if err != nil {
		return err
	}
	defer lr.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, lr)
	return err
}

// LzwReader reads LZW-compressed data from a file, closing both the LZW
// stream and the underlying file on Close.
type LzwReader struct {
	io.ReadCloser
	f *os.File
}

// OpenLzw opens the LZW-compressed file at the specified path for reading,
// using DefaultLzwOrder and DefaultLzwLitWidth.
func OpenLzw(path string) (*LzwReader, error) {
	return OpenLzwWith(path, DefaultLzwOrder, DefaultLzwLitWidth)
}

// OpenLzwWith is like OpenLzw but allows the bit ordering and literal code
// width to be specified explicitly. They must match the values used to
// compress the file.
func OpenLzwWith(path string, order lzw.Order, litWidth int) (*LzwReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &LzwReader{lzw.NewReader(f, order, litWidth), f}, nil
}

// Close closes the LZW stream and the underlying file.
func (r *LzwReader) Close() error {
	if err := r.ReadCloser.Close(); err != nil {
		r.f.Close()
		return err
	}
	return r.f.Close()
}

// LzwWriter writes LZW-compressed data to a file, closing both the LZW
// stream and the underlying file on Close.
type LzwWriter struct {
	io.WriteCloser
	f *os.File
}

// CreateLzw creates a new LZW-compressed file at the specified path for
// writing, using DefaultLzwOrder and DefaultLzwLitWidth.
func CreateLzw(path string) (*LzwWriter, error) {
	return CreateLzwWith(path, DefaultLzwOrder, DefaultLzwLitWidth)
}

// CreateLzwWith is like CreateLzw but allows the bit ordering and literal
// code width to be specified explicitly. litWidth must be in the range [2,8].
func CreateLzwWith(path string, order lzw.Order, litWidth int) (*LzwWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &LzwWriter{lzw.NewWriter(f, order, litWidth), f}, nil
}

// Close flushes and closes the LZW stream, then closes the underlying file.
func (w *LzwWriter) Close() error {
	if err := w.WriteCloser.Close(); err != nil {
		w.f.Close()
		return err
	}
	return w.f.Close()
}

// ----------------------------------------------------------------------------
// BZIP2
// ----------------------------------------------------------------------------
//
// The standard library only implements a bzip2 reader, not a writer, so this
// package cannot offer a Bzip2 compress function without an external
// dependency. Only decompression is supported.

// Bunzip2 decompresses the bzip2-compressed source file into the destination
// file.
func Bunzip2(src, dest string) error {
	br, err := OpenBzip2(src)
	if err != nil {
		return err
	}
	defer br.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, br)
	return err
}

// Bzip2Reader reads bzip2-compressed data from a file, closing the underlying
// file on Close.
type Bzip2Reader struct {
	io.Reader
	f *os.File
}

// OpenBzip2 opens the bzip2-compressed file at the specified path for
// reading.
func OpenBzip2(path string) (*Bzip2Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &Bzip2Reader{bzip2.NewReader(f), f}, nil
}

// Close closes the underlying file.
func (r *Bzip2Reader) Close() error {
	return r.f.Close()
}

// ----------------------------------------------------------------------------
// DECOMPRESS
// ----------------------------------------------------------------------------

// Decompress inspects the source file's contents and extracts or decompresses
// it into dest using whichever of Unzip, Untar, UntarGz, Ungzip, Unzlib or
// Bunzip2 matches. If the format cannot be identified, it returns
// ErrUnknownFormat.
//
// dest is a directory for archive formats (zip, tar, tar.gz) and a file for
// single-stream formats (gzip, zlib, bzip2), matching the corresponding
// UnX function.
//
// Raw DEFLATE (Flate) and LZW streams carry no header identifying them and
// cannot be auto-detected; call Unflate or Unlzw directly for those.
func Decompress(src, dest string) error {
	format, err := detectFormat(src)
	if err != nil {
		return err
	}

	switch format {
	case formatZip:
		return Unzip(src, dest)
	case formatTarGz:
		return UntarGz(src, dest)
	case formatTar:
		return Untar(src, dest)
	case formatGzip:
		return Ungzip(src, dest)
	case formatZlib:
		return Unzlib(src, dest)
	case formatBzip2:
		return Bunzip2(src, dest)
	default:
		return ErrUnknownFormat
	}
}

// ----------------------------------------------------------------------------
// HELPERS
// ----------------------------------------------------------------------------

// archiveFormat identifies a detected archive or compression format.
type archiveFormat int

const (
	formatUnknown archiveFormat = iota
	formatZip
	formatTar
	formatTarGz
	formatGzip
	formatZlib
	formatBzip2
)

// tarMagicOffset is the byte offset of the "ustar" magic within a tar header.
const tarMagicOffset = 257

// relArchiveName computes the name an entry should have inside an archive
// rooted at src: paths relative to src, with src itself (when it is a
// directory) omitted. It returns an empty name for the root directory entry.
func relArchiveName(src, path string, fi os.FileInfo) (string, error) {
	rel, err := filepath.Rel(src, path)
	if err != nil {
		return "", err
	}
	if rel == "." {
		if fi.IsDir() {
			return "", nil
		}
		return filepath.ToSlash(fi.Name()), nil
	}
	return filepath.ToSlash(rel), nil
}

// safeJoin joins dest and name, rejecting names that would escape dest via
// ".." path traversal (a "zip slip").
func safeJoin(dest, name string) (string, error) {
	path := filepath.Join(dest, name)
	cleanDest := filepath.Clean(dest)
	if path != cleanDest && !strings.HasPrefix(path, cleanDest+string(os.PathSeparator)) {
		return "", fmt.Errorf("fsx: illegal file path %q", name)
	}
	return path, nil
}

// extractZipEntry extracts a single zip file entry into dest.
func extractZipEntry(f *zip.File, dest string) error {
	path, err := safeJoin(dest, f.Name)
	if err != nil {
		return err
	}

	if f.FileInfo().IsDir() {
		return os.MkdirAll(path, f.Mode())
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, rc)
	return err
}

// writeTarEntries walks src, writing each file or directory into tw using
// paths relative to src.
func writeTarEntries(tw *tar.Writer, src string) error {
	return filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name, err := relArchiveName(src, path, fi)
		if err != nil {
			return err
		}
		if name == "" {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		header.Name = name

		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tw, file)
		return err
	})
}

// extractTarEntries reads every entry from tr and writes it into dest.
func extractTarEntries(tr *tar.Reader, dest string) error {
	for {
		header, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		path, err := safeJoin(dest, header.Name)
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return err
			}
			out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(out, tr)
			out.Close()
			if err != nil {
				return err
			}
		}
	}
}

// hasTarMagic reports whether the next bytes read from r are a tar header,
// identified by the "ustar" magic at tarMagicOffset.
func hasTarMagic(r io.Reader) bool {
	buf := make([]byte, tarMagicOffset+5)
	n, _ := io.ReadFull(r, buf)
	if n < len(buf) {
		return false
	}
	return string(buf[tarMagicOffset:tarMagicOffset+5]) == "ustar"
}

// isZlibHeader reports whether cmf and flg form a valid zlib header per
// RFC 1950: CM must be 8 (deflate) and (CMF*256+FLG) must be a multiple of 31.
func isZlibHeader(cmf, flg byte) bool {
	if cmf&0x0f != 0x08 {
		return false
	}
	return (uint16(cmf)*256+uint16(flg))%31 == 0
}

// detectFormat sniffs the header of the file at src to identify its archive
// or compression format.
func detectFormat(src string) (archiveFormat, error) {
	f, err := os.Open(src)
	if err != nil {
		return formatUnknown, err
	}
	defer f.Close()

	header := make([]byte, 4)
	n, _ := io.ReadFull(f, header)

	switch {
	case n >= 4 && header[0] == 'P' && header[1] == 'K' &&
		(header[2] == 0x03 || header[2] == 0x05 || header[2] == 0x07) &&
		(header[3] == 0x04 || header[3] == 0x06 || header[3] == 0x08):
		return formatZip, nil

	case n >= 3 && header[0] == 'B' && header[1] == 'Z' && header[2] == 'h':
		return formatBzip2, nil

	case n >= 2 && header[0] == 0x1f && header[1] == 0x8b:
		if _, err := f.Seek(0, io.SeekStart); err != nil {
			return formatUnknown, err
		}
		gz, err := gzip.NewReader(f)
		if err != nil {
			return formatUnknown, err
		}
		defer gz.Close()
		if hasTarMagic(gz) {
			return formatTarGz, nil
		}
		return formatGzip, nil

	case n >= 2 && isZlibHeader(header[0], header[1]):
		return formatZlib, nil
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return formatUnknown, err
	}
	if hasTarMagic(f) {
		return formatTar, nil
	}

	return formatUnknown, nil
}
