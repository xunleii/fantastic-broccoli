package notification

import (
	"testing"
	"fantastic-broccoli/utils"
)

// -- Unit test

func TestBuilderFrom(t *testing.T) {
	b := Builder{}
	b.From("Origin")

	utils.AssertEquals(t, Origin("Origin"), b.from)
}

func TestBuilderTo(t *testing.T) {
	b := Builder{}
	b.To("Destination")

	utils.AssertEquals(t, Destination("Destination"), b.to)
}

func TestBuilderWith(t *testing.T) {
	b := Builder{}
	o := struct {
		a string
		b int
		c bool
	}{"a", 0, false}
	b.With(o)

	utils.AssertEquals(t, o, b.content)
}

func TestBuilderBuild(t *testing.T) {
	b := Builder{}
	o := struct {
		a string
		b int
		c bool
	}{"a", 0, false}
	d1 := Notification{"Origin", "Destination", o}
	d2 := Notification{"", "Destination", o}
	d3 := Notification{"Destination", "Origin", o}
	d4 := Notification{"Destination", "Origin", struct{}{}}

	n := b.From("Origin").To("Destination").With(o).Build()
	utils.AssertEquals(t, d1, n)

	n = b.From("").Build()
	utils.AssertEquals(t, d2, n)

	n = b.From("Destination").To("Origin").Build()
	utils.AssertEquals(t, d3, n)

	n = b.With(struct{}{}).Build()
	utils.AssertEquals(t, d4, n)
}

// -- Benchmark

func BenchmarkBuilderAll(b *testing.B) {
	bd := Builder{}
	ori := Origin("Origin")
	des := Destination("Destination")
	o := struct {
		a string
		b int
		c bool
	}{"a", 0, false}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := bd.From(ori).To(des).With(o).Build()
		n.Content()
	}
}

func BenchmarkBuilderBuild(b *testing.B) {
	bd := Builder{}
	bd.From("Origin").To("Destination").With(struct {
		a string
		b int
		c bool
	}{"a", 0, false})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := bd.Build()
		n.Content()
	}
}

func BenchmarkBuilderNotificationOnly(b *testing.B) {
	ori := Origin("Origin")
	des := Destination("Destination")
	o := struct {
		a string
		b int
		c bool
	}{"a", 0, false}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := Notification{ori, des, o}
		n.Content()
	}
}
