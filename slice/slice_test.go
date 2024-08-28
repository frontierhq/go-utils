package slice

import (
	"reflect"
	"testing"
)

type Driveable interface {
	Drive()
}

type Chargeable interface {
	Charge()
}

type Car struct{}

func (c Car) Drive() {}

type ElectricCar struct {
	Car
}

func (c ElectricCar) Charge() {}

type Motorbike struct{}

func (b Motorbike) Drive() {}

type ElectricMotorbike struct {
	Motorbike
}

func (b ElectricMotorbike) Charge() {}

func TestConvertAllConvertsAnyToStrings(t *testing.T) {
	input := []any{"1", "2", "3"}
	actual, err := ConvertAll[string](input)
	want := "[]string"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsAnyToInts(t *testing.T) {
	input := []any{1, 2, 3}
	actual, err := ConvertAll[int](input)
	want := "[]int"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsAnyToFloats(t *testing.T) {
	input := []any{1.0, 2.0, 3.0}
	actual, err := ConvertAll[float64](input)
	want := "[]float64"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsAnyToInterfaces(t *testing.T) {
	input := []any{Car{}, Motorbike{}, ElectricCar{}}
	actual, err := ConvertAll[Driveable](input)
	want := "[]slice.Driveable"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsStructsToInterfaces(t *testing.T) {
	input := []Car{{}, {}, {}}
	actual, err := ConvertAll[Driveable](input)
	want := "[]slice.Driveable"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsInterfacesToStructs(t *testing.T) {
	input := []Driveable{Car{}, Car{}, Car{}}
	actual, err := ConvertAll[Car](input)
	want := "[]slice.Car"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllConvertsInterfacesToInterfaces(t *testing.T) {
	input := []Driveable{ElectricCar{}, ElectricMotorbike{}, ElectricCar{}}
	actual, err := ConvertAll[Chargeable](input)
	want := "[]slice.Chargeable"
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(input) {
		t.Fatalf("Expected %d results, got %d", len(actual), len(input))
	}
	if reflect.TypeOf(actual).String() != want {
		t.Fatalf("ConvertAll() = %+v, wanted %+v", reflect.TypeOf(actual).String(), want)
	}
}

func TestConvertAllReturnsNilResultErrorWhenConversionFails(t *testing.T) {
	input := []any{"1", "2", 3}
	actual, err := ConvertAll[string](input)
	if err == nil {
		t.Fatal("ConvertAll() should have failed with error")
	}
	if actual != nil {
		t.Fatalf("Expected error nil, got '%+v'", actual)
	}
}

func TestConvertAllFailsWithErrorWhenConvertingIncompatibleBuiltinTypes(t *testing.T) {
	input := []any{"1", "2", 3}
	_, err := ConvertAll[string](input)
	want := "failed to convert source item (sourceType=int, targetType=string)"
	if err == nil {
		t.Fatal("ConvertAll() should have failed with error")
	}
	if err.Error() != want {
		t.Fatalf("Expected error '%s', got '%s'", want, err.Error())
	}
}

func TestConvertAllFailsWithErrorWhenConvertingUnsupportedInterfaceTypes(t *testing.T) {
	input := []Driveable{Car{}, ElectricMotorbike{}, ElectricCar{}}
	_, err := ConvertAll[Chargeable](input)
	want := "failed to convert source item (sourceType=slice.Car, targetType=<nil>)"
	if err == nil {
		t.Fatal("ConvertAll() should have failed with error")
	}
	if err.Error() != want {
		t.Fatalf("Expected error '%s', got '%s'", want, err.Error())
	}
}
