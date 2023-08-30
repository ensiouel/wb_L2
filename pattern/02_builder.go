package main

import "fmt"

/*

Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Плюсы:
- Позволяет создавать продукты пошагово.
- Позволяет использовать один и тот же код для создания различных продуктов.
- Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
- Усложняет код программы из-за введения дополнительных классов.
- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.

*/

func main() {
	redVelvetBuilder := NewRedVelvetBuilder()
	confectioner := NewConfectioner(redVelvetBuilder)

	fmt.Println(confectioner.BuildCake())
}

type CakeBuilder interface {
	Build() Cake
	WithShape() CakeBuilder
	WithTopping() CakeBuilder
	WithDough() CakeBuilder
	WithCoating() CakeBuilder
	WithDiameter() CakeBuilder
	WithLayerCount() CakeBuilder
}

type Confectioner struct {
	cakeBuilder CakeBuilder
}

func NewConfectioner(cakeBuilder CakeBuilder) *Confectioner {
	return &Confectioner{cakeBuilder: cakeBuilder}
}

func (confectioner *Confectioner) BuildCake() Cake {
	return confectioner.cakeBuilder.
		WithShape().
		WithTopping().
		WithDough().
		WithCoating().
		WithDiameter().
		WithLayerCount().
		Build()
}

type Cake struct {
	Shape      string
	Topping    string
	Dough      string
	Coating    string
	Diameter   int
	LayerCount int
}

type RedVelvetBuilder struct {
	Cake
}

func NewRedVelvetBuilder() *RedVelvetBuilder {
	return &RedVelvetBuilder{}
}

func (builder *RedVelvetBuilder) Build() Cake {
	return Cake{
		Shape:      builder.Shape,
		Topping:    builder.Topping,
		Dough:      builder.Dough,
		Coating:    builder.Coating,
		Diameter:   builder.Diameter,
		LayerCount: builder.LayerCount,
	}
}

func (builder *RedVelvetBuilder) WithShape() CakeBuilder {
	builder.Shape = "circle"
	return builder
}

func (builder *RedVelvetBuilder) WithTopping() CakeBuilder {
	builder.Topping = "cream cheese"
	return builder
}

func (builder *RedVelvetBuilder) WithDough() CakeBuilder {
	builder.Dough = "red sponge"
	return builder
}

func (builder *RedVelvetBuilder) WithCoating() CakeBuilder {
	builder.Coating = "white cream"
	return builder
}

func (builder *RedVelvetBuilder) WithDiameter() CakeBuilder {
	builder.Diameter = 16
	return builder
}

func (builder *RedVelvetBuilder) WithLayerCount() CakeBuilder {
	builder.LayerCount = 6
	return builder
}
