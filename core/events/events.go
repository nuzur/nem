package events

func New(params Params) *Implementation {
	return &Implementation{
		client:     params.Client,
		monitoring: params.Monitoring,
		provider:   params.Provider,
	}
}
