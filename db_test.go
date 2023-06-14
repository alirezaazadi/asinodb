package asinodb_test

import (
	"github.com/alirezaazadi/asinodb.git"
	"testing"
)

func TestDatabase_Get(t *testing.T) {

	t.Run("Nothing", func(t *testing.T) {
		d := asinodb.New()

		if _, err := d.Get("key"); err != asinodb.ErrNothing {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("Something", func(t *testing.T) {
		d := asinodb.New()
		d.Set("key", "value")
		v, err := d.Get("key")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if value, ok := v.(string); !ok || value != "value" {
			t.Fatalf("unexpected value %v", v)

		}
	})

}

func TestDatabase_Set(t *testing.T) {
	d := asinodb.New()

	err := d.Set("key", "value")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	v, _ := d.Get("key")

	if value, ok := v.(string); !ok || value != "value" {
		t.Fatalf("unexpected value: %v", v)
	}
}
