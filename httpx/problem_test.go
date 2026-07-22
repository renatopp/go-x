package httpx

import (
	"encoding/json"
	"testing"
)

func TestNewProblem(t *testing.T) {
	p := NewProblem(404, "Not Found", "the resource does not exist")

	if p.Type != "about:blank" {
		t.Errorf("expected Type %q, got %q", "about:blank", p.Type)
	}
	if p.Title != "Not Found" {
		t.Errorf("expected Title %q, got %q", "Not Found", p.Title)
	}
	if p.Status != 404 {
		t.Errorf("expected Status %d, got %d", 404, p.Status)
	}
	if p.Detail != "the resource does not exist" {
		t.Errorf("expected Detail %q, got %q", "the resource does not exist", p.Detail)
	}
	if p.Instance != "" {
		t.Errorf("expected empty Instance, got %q", p.Instance)
	}
}

func TestProblemError(t *testing.T) {
	p := Problem{Title: "Bad Request"}
	if got := p.Error(); got != "Bad Request" {
		t.Errorf("expected Error() %q, got %q", "Bad Request", got)
	}

	var err error = &p
	if got := err.Error(); got != "Bad Request" {
		t.Errorf("expected Error() via error interface %q, got %q", "Bad Request", got)
	}
}

func TestProblemWithType(t *testing.T) {
	p := NewProblem(400, "Bad Request", "")
	got := p.WithType("https://example.com/probs/bad-request")

	if got != p {
		t.Error("expected WithType to return the same pointer")
	}
	if p.Type != "https://example.com/probs/bad-request" {
		t.Errorf("expected Type %q, got %q", "https://example.com/probs/bad-request", p.Type)
	}
}

func TestProblemWithTitle(t *testing.T) {
	p := NewProblem(400, "Bad Request", "")
	got := p.WithTitle("Invalid Request")

	if got != p {
		t.Error("expected WithTitle to return the same pointer")
	}
	if p.Title != "Invalid Request" {
		t.Errorf("expected Title %q, got %q", "Invalid Request", p.Title)
	}
}

func TestProblemWithStatus(t *testing.T) {
	p := NewProblem(400, "Bad Request", "")
	got := p.WithStatus(422)

	if got != p {
		t.Error("expected WithStatus to return the same pointer")
	}
	if p.Status != 422 {
		t.Errorf("expected Status %d, got %d", 422, p.Status)
	}
}

func TestProblemWithDetail(t *testing.T) {
	p := NewProblem(400, "Bad Request", "")
	got := p.WithDetail("field 'name' is required")

	if got != p {
		t.Error("expected WithDetail to return the same pointer")
	}
	if p.Detail != "field 'name' is required" {
		t.Errorf("expected Detail %q, got %q", "field 'name' is required", p.Detail)
	}
}

func TestProblemWithInstance(t *testing.T) {
	p := NewProblem(400, "Bad Request", "")
	got := p.WithInstance("/requests/123")

	if got != p {
		t.Error("expected WithInstance to return the same pointer")
	}
	if p.Instance != "/requests/123" {
		t.Errorf("expected Instance %q, got %q", "/requests/123", p.Instance)
	}
}

func TestProblemWithersChain(t *testing.T) {
	p := NewProblem(500, "Internal Server Error", "unexpected failure").
		WithType("https://example.com/probs/internal").
		WithTitle("Server Error").
		WithStatus(503).
		WithDetail("service unavailable").
		WithInstance("/requests/456")

	if p.Type != "https://example.com/probs/internal" {
		t.Errorf("expected Type %q, got %q", "https://example.com/probs/internal", p.Type)
	}
	if p.Title != "Server Error" {
		t.Errorf("expected Title %q, got %q", "Server Error", p.Title)
	}
	if p.Status != 503 {
		t.Errorf("expected Status %d, got %d", 503, p.Status)
	}
	if p.Detail != "service unavailable" {
		t.Errorf("expected Detail %q, got %q", "service unavailable", p.Detail)
	}
	if p.Instance != "/requests/456" {
		t.Errorf("expected Instance %q, got %q", "/requests/456", p.Instance)
	}
}

func TestProblemJSONMarshalling(t *testing.T) {
	p := NewProblem(400, "Bad Request", "field 'name' is required").WithInstance("/requests/123")

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("unexpected error marshalling: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}

	if decoded["type"] != "about:blank" {
		t.Errorf("expected type %q, got %v", "about:blank", decoded["type"])
	}
	if decoded["title"] != "Bad Request" {
		t.Errorf("expected title %q, got %v", "Bad Request", decoded["title"])
	}
	if decoded["status"] != float64(400) {
		t.Errorf("expected status %v, got %v", 400, decoded["status"])
	}
	if decoded["detail"] != "field 'name' is required" {
		t.Errorf("expected detail %q, got %v", "field 'name' is required", decoded["detail"])
	}
	if decoded["instance"] != "/requests/123" {
		t.Errorf("expected instance %q, got %v", "/requests/123", decoded["instance"])
	}
}

func TestProblemJSONOmitsEmptyOptionalFields(t *testing.T) {
	p := NewProblem(404, "Not Found", "")

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("unexpected error marshalling: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}

	if _, ok := decoded["detail"]; ok {
		t.Errorf("expected detail to be omitted when empty, got %v", decoded["detail"])
	}
	if _, ok := decoded["instance"]; ok {
		t.Errorf("expected instance to be omitted when empty, got %v", decoded["instance"])
	}
}
