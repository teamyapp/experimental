package channel

type ChannelFactory interface {
	NewChannel() Channel
}

type ChannelFactoryImpl struct {

}


func (c ChannelFactoryImpl) NewChannel() Channel {
	panic("implement me")
}

var _ ChannelFactory = (*ChannelFactoryImpl)(nil)

