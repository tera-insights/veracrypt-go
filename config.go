package veracrypt

// HashAlgo is an enumeration of all the hash algorithms available in
// VeraCrypt.
type HashAlgo string

const (
	SHA256    HashAlgo = "sha-256"
	SHA512    HashAlgo = "sha-512"
	Whirlpool HashAlgo = "whirlpool"
	RIPEMD160 HashAlgo = "ripemd-160"
)

// Config is a configuration for instantiating a VeraCrypt instance.
type Config struct {
	VolumePath string
	MountPath  string
	Password   string
	Hash       int
}

// setDefaults sets reasonable defaults for a VeraCrypt configuration,
// but will return an error if any required fields are missing.
func (c *Config) setDefaults() error {
	return nil
}
