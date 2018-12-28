package config;

import (
	"testing"
)

func Test_RssParser_Rules(t *testing.T) {
	cfg, err := GetConfig()
	if err != nil {
		t.Error(err)
	}

	if len(cfg.RSSParser) == 0 {
		t.Error("No RSSParser instances configured")
	}

	if firstTest, ok := cfg.RSSParser["firsttest"]; ok {
		if len(firstTest.Rules) == 0 {
			t.Error("No rules defined in the configuration")
		}

		hebsub := firstTest.Rules[0]

		if hebsub.Pattern != "hebsub" {
			t.Error("Pattern is not 'hebsub'")
		}

		if hebsub.Matching != "title" {
			t.Error("Matching is not 'title'")
		}

		if val, ok := ConvertRuleValueToString(hebsub.All); ok && val != "reject" {
			t.Error("All is not 'reject'")
		}

		internal := firstTest.Rules[1]
		if internal.Pattern != "internal" {
			t.Error("Pattern is not internal")
		}

		if internal.Matching != "flags" {
			t.Error("Matching is not 'flags'")
		}

		if val,ok := ConvertRuleValueToFloat(internal.First); ok && val != 0.8 {
			t.Error("All is not '0.8'")
		}

		if val,ok := ConvertRuleValueToFloat(internal.Other); ok && val != 0.001 {
			t.Error("Other is not '0.001'")
		}
		_,ok := ConvertRuleValueToString(internal.Other)
		t.Error(ok)

	} else {
		t.Error("No key 'firsttest' found in configuration")
	}
}