package utils

type Task struct {
	Config         map[string]string
	CrackingEngine *Process
	Generator      *Process
}

func NewTask(taskConfig map[string]string, crackingEnginePath string, generatorPath string) *Task {
	return &Task{
		Config:         taskConfig,
		CrackingEngine: NewProcess(crackingEnginePath),
		Generator:      NewProcess(generatorPath),
	}
}
