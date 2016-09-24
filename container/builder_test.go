package container

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

type Foo struct{}
type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

func TestRegisterWithInstance(t *testing.T) {
	Convey("Given some instance of an arbitrary type", t, func() {
		expected := new(Foo)

		Convey("When the instance is registered into the container builder ", func() {
			builder := NewBuilder()
			b, err := builder.Register("foo", expected)

			Convey("Then the register method should return the builder", func() {
				So(err, ShouldBeEmpty)
			})

			Convey("And the register method should return an empty error", func() {
				So(b, ShouldEqual, b)
			})
		})
	})
}

func TestRegisterWithConstructorFunction(t *testing.T) {
	Convey("Given some constructor function that returns an arbitrary type", t, func() {
		Convey("When the function is registered into the container builder ", func() {
			builder := NewBuilder()
			b, err := builder.Register("foo", NewBar)

			Convey("Then the register method should return the builder", func() {
				So(err, ShouldBeEmpty)
			})

			Convey("And the register method should return an empty error", func() {
				So(b, ShouldEqual, b)
			})
		})
	})
}

func TestRegisterHasDefinition(t *testing.T) {
	builder := NewBuilder()
	builder.Register("foo", new(Foo))

	assert.True(t, builder.HasDefinition("foo"))
	assert.False(t, builder.HasDefinition("bar"))
}

func TestRegisterGetDefinitionRegisteredWithInstance(t *testing.T) {
	builder := NewBuilder()
	actual := new(Foo)
	builder.Register("foo", actual)

	foo, err := builder.Get("foo")

	if err != nil {
		t.Log(fmt.Sprintf("%s", err))
		t.Fail()
	}

	var expected *Foo
	expected = foo.(*Foo)

	assert.Nil(t, expected)
	assert.IsType(t, expected, actual)
	assert.ObjectsAreEqual(expected, actual)
}

func TestRegisterGetDefinitionRegisteredWithConstructorFunction(t *testing.T) {
	builder := NewBuilder()
	builder.Register("bar", NewBar)

	bar, err := builder.Get("bar")

	if err != nil {
		t.Log(fmt.Sprintf("%s", err))
		t.Fail()
	}

	var expected *Bar
	expected = bar.(*Bar)

	assert.Nil(t, expected)
	assert.IsType(t, expected, bar)
	assert.ObjectsAreEqual(expected, bar)
}

func TestRegisterReturnsErrorWhenDefinitionAlreadyExist(t *testing.T) {
	builder := NewBuilder()

	builder.Register("bar", NewBar)
	_, err := builder.Register("bar", NewBar)

	assert.Error(t, err)
}

func TestRegisterReturnsErrorWhenServiceIsNotFound(t *testing.T) {
	builder := NewBuilder()
	_, err := builder.Get("service_that_does_not_exist")

	assert.Error(t, err)
}
