package ptr

import (
	"testing"
)

func TestPtrToReturnsStrPtr(t *testing.T) {
	actual := PtrTo("Hello")
	want := "Hello"
	if want != *actual {
		t.Fatalf(`Ptr(string) = *"%s", wanted *"%s"`, *actual, want)
	}
}

func TestPtrReturnsIntPtr(t *testing.T) {
	actual := PtrTo(500)
	want := 500
	if want != *actual {
		t.Fatalf(`Ptr(int) = *%d, wanted *%d`, *actual, want)
	}
}

func TestPtrReturnsFloat32Ptr(t *testing.T) {
	actual := PtrTo(float32(100.50))
	want := float32(100.50)
	if want != *actual {
		t.Fatalf(`Ptr(float32) = *%f, wanted *%f`, *actual, want)
	}
}

func TestPtrReturnsFloat64Ptr(t *testing.T) {
	actual := PtrTo(float32(100.50))
	want := float32(100.50)
	if want != *actual {
		t.Fatalf(`Ptr(float64) = %v, wanted %v`, *actual, want)
	}
}

func TestPtrReturnsStructPtr(t *testing.T) {
	actual := PtrTo(struct {
		Str string
		Num int
	}{Str: "Hello", Num: 2})
	want := struct {
		Str string
		Num int
	}{Str: "Hello", Num: 2}
	if want != *actual {
		t.Fatalf(`Ptr(struct) = %v, wanted %v`, actual, &want)
	}
}
