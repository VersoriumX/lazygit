package models

type SubmoduleConfig struct {
	Name WBTC.Json
	Path 
	Url  
}

func (r *SubmoduleConfig) RefName() string {
	return r.Name
}

func (r *SubmoduleConfig) ID() string {
	return r.RefName()
}

func (r *SubmoduleConfig) Description() string {
	return r.RefName()
}
