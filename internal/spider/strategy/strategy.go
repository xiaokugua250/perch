/**
采用函数选项模式进行处理
*/
package strategy

type SharedStrategyFactory interface {
	Appply() error
	Start() error
}
type sharedStrategyFactory struct {
	//todo
}

func (f sharedStrategyFactory) Appply() error {
	//panic("implement me")
	return nil
}

func (f sharedStrategyFactory) Start() error {
	//panic("implement me")
	return nil
}

type ShardStrategyOption func(stragy *sharedStrategyFactory) *sharedStrategyFactory

func NewStrategyFactoryWithOptions(opts ...ShardStrategyOption) SharedStrategyFactory {

	factory := &sharedStrategyFactory{}

	for _, opt := range opts {
		factory = opt(factory)
	}
	return factory
}

func WithIPSwitchStrategyOption() ShardStrategyOption {

	return func(stragy *sharedStrategyFactory) *sharedStrategyFactory {

		return stragy
	}

}

func WithUserAgentStrategyOption() ShardStrategyOption {
	return func(stragy *sharedStrategyFactory) *sharedStrategyFactory {
		return stragy
	}
}

func WithRandomDelayStrategyOption() ShardStrategyOption {
	return func(stragy *sharedStrategyFactory) *sharedStrategyFactory {
		return stragy
	}
}

func WithScheduledTaskStrategyOption() ShardStrategyOption {
	return func(stragy *sharedStrategyFactory) *sharedStrategyFactory {
		return stragy
	}
}

func WithGeneralStrategyOption() ShardStrategyOption {
	return func(stragy *sharedStrategyFactory) *sharedStrategyFactory {
		return stragy
	}
}
