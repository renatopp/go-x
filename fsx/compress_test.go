package fsx_test

import (
	"archive/tar"
	"archive/zip"
	"compress/lzw"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/renatopp/go-x/fsx"
	"github.com/renatopp/go-x/testx"
)

// bzip2Fixture is the bzip2 compression of "hello fsx bzip2 test fixture\n",
// produced with the system bzip2 tool. The standard library has no bzip2
// writer, so a fixture is needed to test Bunzip2 and Decompress.
var bzip2Fixture = []byte{
	0x42, 0x5a, 0x68, 0x39, 0x31, 0x41, 0x59, 0x26, 0x53, 0x59, 0xc3, 0x15,
	0x89, 0x66, 0x00, 0x00, 0x06, 0x59, 0x80, 0x00, 0x10, 0x40, 0x00, 0x10,
	0x00, 0x13, 0x64, 0xde, 0x50, 0x20, 0x00, 0x31, 0x43, 0x4d, 0x30, 0x00,
	0x44, 0x32, 0x64, 0x30, 0x6a, 0x25, 0xd3, 0xac, 0x34, 0xba, 0x99, 0x74,
	0x43, 0xd1, 0x50, 0x34, 0x1c, 0xbf, 0x1b, 0x0f, 0x8b, 0xb9, 0x22, 0x9c,
	0x28, 0x48, 0x61, 0x8a, 0xc4, 0xb3, 0x00,
}

const bzip2FixtureContent = "hello fsx bzip2 test fixture\n"

func setupSrcDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	testx.Nil(t, os.WriteFile(filepath.Join(dir, "a.txt"), []byte("file a"), 0644))
	testx.Nil(t, os.MkdirAll(filepath.Join(dir, "sub"), 0755))
	testx.Nil(t, os.WriteFile(filepath.Join(dir, "sub", "b.txt"), []byte("file b"), 0644))
	return dir
}

func verifyExtractedDir(t *testing.T, dir string) {
	t.Helper()
	a, err := os.ReadFile(filepath.Join(dir, "a.txt"))
	testx.Nil(t, err)
	testx.Equal(t, "file a", string(a))

	b, err := os.ReadFile(filepath.Join(dir, "sub", "b.txt"))
	testx.Nil(t, err)
	testx.Equal(t, "file b", string(b))
}

// ----------------------------------------------------------------------------
// ZIP
// ----------------------------------------------------------------------------

func TestZipUnzipDir(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.zip")
	testx.Nil(t, fsx.Zip(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Unzip(archive, dest))
	verifyExtractedDir(t, dest)
}

func TestZipUnzipFile(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("solo"), 0644))

	archive := filepath.Join(dir, "out.zip")
	testx.Nil(t, fsx.Zip(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Unzip(archive, dest))

	data, err := os.ReadFile(filepath.Join(dest, "a.txt"))
	testx.Nil(t, err)
	testx.Equal(t, "solo", string(data))
}

func TestUnzipRejectsPathTraversal(t *testing.T) {
	dir := t.TempDir()
	archive := filepath.Join(dir, "evil.zip")

	f, err := os.Create(archive)
	testx.Nil(t, err)
	zw := zip.NewWriter(f)
	w, err := zw.Create("../evil.txt")
	testx.Nil(t, err)
	_, err = w.Write([]byte("pwned"))
	testx.Nil(t, err)
	testx.Nil(t, zw.Close())
	testx.Nil(t, f.Close())

	err = fsx.Unzip(archive, t.TempDir())
	testx.NotNil(t, err)
}

func TestOpenCreateZip(t *testing.T) {
	dir := t.TempDir()
	archive := filepath.Join(dir, "raw.zip")

	zw, err := fsx.CreateZip(archive)
	testx.Nil(t, err)
	w, err := zw.Create("a.txt")
	testx.Nil(t, err)
	_, err = w.Write([]byte("raw"))
	testx.Nil(t, err)
	testx.Nil(t, zw.Close())

	zr, err := fsx.OpenZip(archive)
	testx.Nil(t, err)
	defer zr.Close()
	testx.Equal(t, 1, len(zr.File))
}

// ----------------------------------------------------------------------------
// TAR
// ----------------------------------------------------------------------------

func TestTarUntarDir(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.tar")
	testx.Nil(t, fsx.Tar(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Untar(archive, dest))
	verifyExtractedDir(t, dest)
}

func TestTarUntarFile(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("solo"), 0644))

	archive := filepath.Join(dir, "out.tar")
	testx.Nil(t, fsx.Tar(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Untar(archive, dest))

	data, err := os.ReadFile(filepath.Join(dest, "a.txt"))
	testx.Nil(t, err)
	testx.Equal(t, "solo", string(data))
}

func TestUntarRejectsPathTraversal(t *testing.T) {
	dir := t.TempDir()
	archive := filepath.Join(dir, "evil.tar")

	f, err := os.Create(archive)
	testx.Nil(t, err)
	tw := tar.NewWriter(f)
	testx.Nil(t, tw.WriteHeader(&tar.Header{Name: "../evil.txt", Size: 5, Mode: 0644}))
	_, err = tw.Write([]byte("pwned"))
	testx.Nil(t, err)
	testx.Nil(t, tw.Close())
	testx.Nil(t, f.Close())

	err = fsx.Untar(archive, t.TempDir())
	testx.NotNil(t, err)
}

func TestOpenCreateTar(t *testing.T) {
	dir := t.TempDir()
	archive := filepath.Join(dir, "raw.tar")

	tw, err := fsx.CreateTar(archive)
	testx.Nil(t, err)
	testx.Nil(t, tw.WriteHeader(&tar.Header{Name: "a.txt", Size: 3, Mode: 0644}))
	_, err = tw.Write([]byte("raw"))
	testx.Nil(t, err)
	testx.Nil(t, tw.Close())

	tr, err := fsx.OpenTar(archive)
	testx.Nil(t, err)
	defer tr.Close()

	h, err := tr.Next()
	testx.Nil(t, err)
	testx.Equal(t, "a.txt", h.Name)
}

// ----------------------------------------------------------------------------
// TAR+GZIP
// ----------------------------------------------------------------------------

func TestTarGzUntarGzDir(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.tar.gz")
	testx.Nil(t, fsx.TarGz(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.UntarGz(archive, dest))
	verifyExtractedDir(t, dest)
}

// ----------------------------------------------------------------------------
// GZIP
// ----------------------------------------------------------------------------

func TestGzipUngzip(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("compress me"), 0644))

	archive := filepath.Join(dir, "a.gz")
	testx.Nil(t, fsx.Gzip(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Ungzip(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "compress me", string(data))
}

func TestGzipDirReturnsErrIsDir(t *testing.T) {
	dir := setupSrcDir(t)
	err := fsx.Gzip(dir, filepath.Join(t.TempDir(), "out.gz"))
	testx.True(t, errors.Is(err, fsx.ErrIsDir))
}

func TestOpenCreateGzip(t *testing.T) {
	dir := t.TempDir()
	archive := filepath.Join(dir, "raw.gz")

	gw, err := fsx.CreateGzip(archive)
	testx.Nil(t, err)
	_, err = gw.Write([]byte("raw"))
	testx.Nil(t, err)
	testx.Nil(t, gw.Close())

	gr, err := fsx.OpenGzip(archive)
	testx.Nil(t, err)
	defer gr.Close()

	data, err := io.ReadAll(gr)
	testx.Nil(t, err)
	testx.Equal(t, "raw", string(data))
}

// ----------------------------------------------------------------------------
// ZLIB
// ----------------------------------------------------------------------------

func TestZlibUnzlib(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("compress me"), 0644))

	archive := filepath.Join(dir, "a.zlib")
	testx.Nil(t, fsx.Zlib(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Unzlib(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "compress me", string(data))
}

func TestZlibDirReturnsErrIsDir(t *testing.T) {
	dir := setupSrcDir(t)
	err := fsx.Zlib(dir, filepath.Join(t.TempDir(), "out.zlib"))
	testx.True(t, errors.Is(err, fsx.ErrIsDir))
}

// ----------------------------------------------------------------------------
// FLATE
// ----------------------------------------------------------------------------

func TestFlateUnflate(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("compress me"), 0644))

	archive := filepath.Join(dir, "a.flate")
	testx.Nil(t, fsx.Flate(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Unflate(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "compress me", string(data))
}

func TestFlateDirReturnsErrIsDir(t *testing.T) {
	dir := setupSrcDir(t)
	err := fsx.Flate(dir, filepath.Join(t.TempDir(), "out.flate"))
	testx.True(t, errors.Is(err, fsx.ErrIsDir))
}

// ----------------------------------------------------------------------------
// LZW
// ----------------------------------------------------------------------------

func TestLzwUnlzw(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("lzwlzwlzwlzw"), 0644))

	archive := filepath.Join(dir, "a.lzw")
	testx.Nil(t, fsx.Lzw(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Unlzw(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "lzwlzwlzwlzw", string(data))
}

func TestLzwWithExplicitParams(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("lzwlzwlzwlzw"), 0644))

	archive := filepath.Join(dir, "a.lzw")
	testx.Nil(t, fsx.LzwWith(src, archive, lzw.LSB, 8))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.UnlzwWith(archive, dest, lzw.LSB, 8))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "lzwlzwlzwlzw", string(data))
}

func TestLzwDirReturnsErrIsDir(t *testing.T) {
	dir := setupSrcDir(t)
	err := fsx.Lzw(dir, filepath.Join(t.TempDir(), "out.lzw"))
	testx.True(t, errors.Is(err, fsx.ErrIsDir))
}

// ----------------------------------------------------------------------------
// BZIP2
// ----------------------------------------------------------------------------

func TestBunzip2(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.bz2")
	testx.Nil(t, os.WriteFile(src, bzip2Fixture, 0644))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Bunzip2(src, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, bzip2FixtureContent, string(data))
}

// ----------------------------------------------------------------------------
// DECOMPRESS
// ----------------------------------------------------------------------------

func TestDecompressZip(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.zip")
	testx.Nil(t, fsx.Zip(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Decompress(archive, dest))
	verifyExtractedDir(t, dest)
}

func TestDecompressTar(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.tar")
	testx.Nil(t, fsx.Tar(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Decompress(archive, dest))
	verifyExtractedDir(t, dest)
}

func TestDecompressTarGz(t *testing.T) {
	src := setupSrcDir(t)
	archive := filepath.Join(t.TempDir(), "out.tar.gz")
	testx.Nil(t, fsx.TarGz(src, archive))

	dest := t.TempDir()
	testx.Nil(t, fsx.Decompress(archive, dest))
	verifyExtractedDir(t, dest)
}

func TestDecompressGzip(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("data"), 0644))

	archive := filepath.Join(dir, "a.gz")
	testx.Nil(t, fsx.Gzip(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Decompress(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "data", string(data))
}

func TestDecompressZlib(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.txt")
	testx.Nil(t, os.WriteFile(src, []byte("data"), 0644))

	archive := filepath.Join(dir, "a.zlib")
	testx.Nil(t, fsx.Zlib(src, archive))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Decompress(archive, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, "data", string(data))
}

func TestDecompressBzip2(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.bz2")
	testx.Nil(t, os.WriteFile(src, bzip2Fixture, 0644))

	dest := filepath.Join(dir, "out.txt")
	testx.Nil(t, fsx.Decompress(src, dest))

	data, err := os.ReadFile(dest)
	testx.Nil(t, err)
	testx.Equal(t, bzip2FixtureContent, string(data))
}

func TestDecompressUnknownFormat(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a.bin")
	testx.Nil(t, os.WriteFile(src, []byte{0x00, 0x01, 0x02, 0x03, 0x04}, 0644))

	err := fsx.Decompress(src, filepath.Join(dir, "out"))
	testx.True(t, errors.Is(err, fsx.ErrUnknownFormat))
}
