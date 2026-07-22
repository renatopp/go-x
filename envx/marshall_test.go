package envx

import (
	"testing"
	"time"
)

type testConfig struct {
	Name    string        `env:"TC_NAME"`
	Port    uint32        `env:"TC_PORT,default=9090"`
	Debug   bool          `env:"TC_DEBUG"`
	Score   float64       `env:"TC_SCORE"`
	Timeout time.Duration `env:"TC_TIMEOUT,default=10s"`
	Workers int           `env:"TC_WORKERS"`
	ignored string        `env:"TC_IGNORED"`
	NoTag   string
}

func TestUnmarshal(t *testing.T) {
	t.Setenv("TC_NAME", "matchmaking")
	t.Setenv("TC_PORT", "8080")
	t.Setenv("TC_DEBUG", "true")
	t.Setenv("TC_SCORE", "9.5")
	t.Setenv("TC_TIMEOUT", "30s")
	t.Setenv("TC_WORKERS", "-4")

	var cfg testConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Name != "matchmaking" {
		t.Errorf("Name: got %q, want %q", cfg.Name, "matchmaking")
	}
	if cfg.Port != 8080 {
		t.Errorf("Port: got %d, want 8080", cfg.Port)
	}
	if !cfg.Debug {
		t.Errorf("Debug: got false, want true")
	}
	if cfg.Score != 9.5 {
		t.Errorf("Score: got %f, want 9.5", cfg.Score)
	}
	if cfg.Timeout != 30*time.Second {
		t.Errorf("Timeout: got %v, want 30s", cfg.Timeout)
	}
	if cfg.Workers != -4 {
		t.Errorf("Workers: got %d, want -4", cfg.Workers)
	}
}

func TestUnmarshal_DefaultValues(t *testing.T) {
	var cfg testConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != 9090 {
		t.Errorf("Port default: got %d, want 9090", cfg.Port)
	}
	if cfg.Timeout != 10*time.Second {
		t.Errorf("Timeout default: got %v, want 10s", cfg.Timeout)
	}
}

func TestUnmarshal_EnvOverridesDefault(t *testing.T) {
	t.Setenv("TC_PORT", "3000")
	t.Setenv("TC_TIMEOUT", "1m")

	var cfg testConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != 3000 {
		t.Errorf("Port override: got %d, want 3000", cfg.Port)
	}
	if cfg.Timeout != time.Minute {
		t.Errorf("Timeout override: got %v, want 1m", cfg.Timeout)
	}
}

func TestUnmarshal_MissingEnvLeavesZeroValue(t *testing.T) {
	var cfg testConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Name != "" || cfg.Workers != 0 {
		t.Errorf("expected zero values, got Name=%q Workers=%d", cfg.Name, cfg.Workers)
	}
}

func TestUnmarshal_InvalidValue(t *testing.T) {
	t.Setenv("TC_PORT", "not-a-number")

	var cfg testConfig
	if err := Unmarshal(&cfg); err == nil {
		t.Fatal("expected error for invalid uint, got nil")
	}
}

func TestUnmarshal_RequiresPointerToStruct(t *testing.T) {
	if err := Unmarshal("not a struct"); err == nil {
		t.Fatal("expected error for non-struct, got nil")
	}
	cfg := testConfig{}
	if err := Unmarshal(cfg); err == nil {
		t.Fatal("expected error for non-pointer, got nil")
	}
}

// ---- required ---------------------------------------------------------------

type requiredConfig struct {
	Token string `env:"TC_TOKEN,required"`
}

func TestUnmarshal_RequiredPresent(t *testing.T) {
	t.Setenv("TC_TOKEN", "secret")

	var cfg requiredConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Token != "secret" {
		t.Errorf("Token: got %q, want %q", cfg.Token, "secret")
	}
}

func TestUnmarshal_RequiredMissing(t *testing.T) {
	var cfg requiredConfig
	if err := Unmarshal(&cfg); err == nil {
		t.Fatal("expected error for missing required field, got nil")
	}
}

// ---- slices -----------------------------------------------------------------

type sliceConfig struct {
	Tags   []string        `env:"TC_TAGS,default=a,b,c"`
	Ports  []uint16        `env:"TC_PORTS"`
	Scores []float64       `env:"TC_SCORES"`
	Delays []time.Duration `env:"TC_DELAYS"`
}

func TestUnmarshal_SliceFromEnv(t *testing.T) {
	t.Setenv("TC_PORTS", "8080,8081,8082")
	t.Setenv("TC_SCORES", "1.1,2.2,3.3")
	t.Setenv("TC_DELAYS", "1s,500ms,2m")

	var cfg sliceConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Ports) != 3 || cfg.Ports[0] != 8080 || cfg.Ports[2] != 8082 {
		t.Errorf("Ports: got %v", cfg.Ports)
	}
	if len(cfg.Scores) != 3 || cfg.Scores[1] != 2.2 {
		t.Errorf("Scores: got %v", cfg.Scores)
	}
	if len(cfg.Delays) != 3 || cfg.Delays[0] != time.Second || cfg.Delays[2] != 2*time.Minute {
		t.Errorf("Delays: got %v", cfg.Delays)
	}
}

func TestUnmarshal_SliceDefault(t *testing.T) {
	var cfg sliceConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Tags) != 3 || cfg.Tags[0] != "a" || cfg.Tags[2] != "c" {
		t.Errorf("Tags default: got %v", cfg.Tags)
	}
}

func TestUnmarshal_SliceEnvOverridesDefault(t *testing.T) {
	t.Setenv("TC_TAGS", "x,y")

	var cfg sliceConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Tags) != 2 || cfg.Tags[0] != "x" || cfg.Tags[1] != "y" {
		t.Errorf("Tags override: got %v", cfg.Tags)
	}
}

// ---- nested structs -----------------------------------------------------

type nestedInner struct {
	Host string `env:"TC_HOST,default=localhost"`
	Port uint32 `env:"TC_INNER_PORT,default=5432"`
}

type nestedConfig struct {
	Name  string `env:"TC_NAME"`
	Inner nestedInner
}

func TestUnmarshal_NestedStruct(t *testing.T) {
	t.Setenv("TC_NAME", "matchmaking")
	t.Setenv("TC_HOST", "db.internal")
	t.Setenv("TC_INNER_PORT", "6543")

	var cfg nestedConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Name != "matchmaking" {
		t.Errorf("Name: got %q, want %q", cfg.Name, "matchmaking")
	}
	if cfg.Inner.Host != "db.internal" {
		t.Errorf("Inner.Host: got %q, want %q", cfg.Inner.Host, "db.internal")
	}
	if cfg.Inner.Port != 6543 {
		t.Errorf("Inner.Port: got %d, want 6543", cfg.Inner.Port)
	}
}

func TestUnmarshal_NestedStructDefaults(t *testing.T) {
	var cfg nestedConfig
	if err := Unmarshal(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Inner.Host != "localhost" {
		t.Errorf("Inner.Host default: got %q, want %q", cfg.Inner.Host, "localhost")
	}
	if cfg.Inner.Port != 5432 {
		t.Errorf("Inner.Port default: got %d, want 5432", cfg.Inner.Port)
	}
}
