## 工厂模式

工厂模式实际就是通过一个工厂创建出类形同但是过程不同的对象

### 简单工厂
比如 最简单的简单工厂
直接返回一个对象
```go
package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// NewIRuleConfigParser NewIRuleConfigParser
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
```

### 工厂方法
工厂方法模式，就是将简单工厂模式的工厂方法抽象出来，通过工厂方法返回一个工厂，通过工厂方法返回一个对象。
```go
// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
```
返回一个`IRuleConfigParserFactory`，通过`IRuleConfigParserFactory`调用`CreateParser()`方法返回一个`IRuleConfigParser对象`

### 抽象工厂
通过定义一个抽象工厂，通过抽象工厂创建一个抽象产品。
这样每个工厂都可以自定义的返回自己实现的抽象产品

这样所有的实现之间都不再直接依赖具体实现，而是依赖于接口(除了工厂生产具体的产品之外)。
在 `simple_pizza_factory`中
`pizzaStory`依赖`pizzaCreator`这个工厂，实现这个工厂的有`beijingPizzaStory`
在`pizzaStory`中，创建pizza的时候就不再需要依赖具体的实例，也不需要知道具体的实例，只需要知道创建了一个pizza即可，
工厂的选择权交给用户的决定

每个不同的披萨店都有相同的产品，比如 `cheesePizza`，`clamPizza`等等，但是因为地区的不通过，每个披萨所使用的原料和口味可能不太相同，
所以我们的每个披萨店都需要自己的披萨原料工厂，即：`pizzaIngredientFactory`，针对这个披萨原料接口，都各自工厂各自不同的实现，
在披萨店创建披萨的时候，选择对应的披萨原料工厂即可，也不需要关注具体的披萨原料实现，只需要知道创建了一个披萨原料工厂即可。

在pizza准备阶段调用该披萨的`prepare`方法， `prepare`方法中调用披萨原料工厂的`createXXX`方法，通过披萨原料工厂创建披萨原料。
因为每个披萨都是实现了 `pizza`接口，并且内置了披萨原料工厂接口。


