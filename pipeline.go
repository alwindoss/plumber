package plumber

// Pipeline is the interface that defines the operations the pipeline can do
type Pipeline interface {
	AddSource(s Sourcer) Pipeline
	AddProcessor(p Processor) Pipeline
	AddSink(s Sinker) Pipeline
	Run()
}

// PipelineBuilder is the interface that defines the operations the pipeline can do
type PipelineBuilder interface {
	Build() Pipeline
}

// NewPipelineBuilder is used to create a Pipeline
func NewPipelineBuilder() Pipeline {
	return &pipeline{}
}

type pipeline struct {
	source     Sourcer
	processors []Processor
	sink       Sinker
}

func (pipe *pipeline) AddSource(s Sourcer) Pipeline {
	pipe.source = s
	return pipe
}

func (pipe *pipeline) AddProcessor(p Processor) Pipeline {
	pipe.processors = append(pipe.processors, p)
	return pipe
}

func (pipe *pipeline) AddSink(s Sinker) Pipeline {
	pipe.sink = s
	return pipe
}

func (pipe *pipeline) Run() {
	sourceOutCh := make(chan Message)
	go pipe.source.Source(sourceOutCh)

	totalProcs := len(pipe.processors)
	var sinkInCh chan Message

	for i, processor := range pipe.processors {
		procOutCh := make(chan Message)
		if i == 0 {
			processor.Process(sourceOutCh, procOutCh)
		} else {
			procInCh := make(chan Message)
			processor.Process(procInCh, procOutCh)
			if i == (totalProcs - 1) {
				sinkInCh = procOutCh
			}
		}
	}

	pipe.sink.Sink(sinkInCh)

}
