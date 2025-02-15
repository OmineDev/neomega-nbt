package caster

import "github.com/mitchellh/mapstructure"

func MapNBT(input any, out any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "nbt",
		Result:  out,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
