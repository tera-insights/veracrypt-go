package veracrypt

// Instance is a controller for a running VeraCrypt process.
type Instance struct {
}

// New starts a new VeraCrypt process.
func New(conf *Config) (inst Instance, err error) {
	if err = conf.setDefaults(); err != nil {
		return
	}
	return
}

func CreateVolume(path string, sizeBytes uint64)
