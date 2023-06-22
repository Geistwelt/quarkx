/**
* Author: Xiangyu Wu
* Date: 2023-06-21
* From: hyperledger/fabric/common/metrics/internal/namer/namer.go
 */

package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/geistwelt/quarkx/common/metrics"
)

type Namer struct {
	namespace  string
	subsystem  string
	name       string
	nameFormat string
	labelNames map[string]struct{}
}

func NewCounterNamer(opt metrics.CounterOpts) *Namer {
	return &Namer{
		namespace:  opt.Namespace,
		subsystem:  opt.Subsystem,
		name:       opt.Name,
		nameFormat: opt.StatsdFormat,
		labelNames: sliceToSet(opt.LabelNames),
	}
}

func NewGaugeNamer(opt metrics.GaugeOpts) *Namer {
	return &Namer{
		namespace:  opt.Namespace,
		subsystem:  opt.Subsystem,
		name:       opt.Name,
		nameFormat: opt.StatsdFormat,
		labelNames: sliceToSet(opt.LabelNames),
	}
}

func NewHistogramNamer(opt metrics.HistogramOpts) *Namer {
	return &Namer{
		namespace:  opt.Namespace,
		subsystem:  opt.Subsystem,
		name:       opt.Name,
		nameFormat: opt.StatsdFormat,
		labelNames: sliceToSet(opt.LabelNames),
	}
}

func (n *Namer) validateKey(name string) {
	if _, ok := n.labelNames[name]; !ok {
		panic("invalid label name: " + name)
	}
}

func (n *Namer) FullyQualifiedName() string {
	switch {
	case n.namespace != "" && n.subsystem != "":
		return strings.Join([]string{n.namespace, n.subsystem, n.name}, ".")
	case n.namespace != "":
		return strings.Join([]string{n.namespace, n.name}, ".")
	case n.subsystem != "":
		return strings.Join([]string{n.subsystem, n.name}, ".")
	default:
		return n.name
	}
}

// 将 {k1, v1, k2, v2, ...} 转换为 {k1:v1, k2:v2, ...}
func (n *Namer) labelsToMap(labelValues []string) map[string]string {
	labels := make(map[string]string)
	for i := 0; i < len(labelValues); i += 2 {
		key := labelValues[i]
		n.validateKey(key)
		if i == len(labelValues)-1 {
			labels[key] = "unknown"
		} else {
			labels[key] = labelValues[i+1]
		}
	}
	return labels
}

var formatRegexp = regexp.MustCompile(`%{([#?[:alnum:]_]+)}`)
var invalidLabelValueRegexp = regexp.MustCompile(`[.|:\s]`)

func (n *Namer) Format(labelValues ...string) string {
	// labelValues 里的每个 key 都必须在 Namer.labelNames 存在。
	labels := n.labelsToMap(labelValues)

	cursor := 0
	var segments []string
	matches := formatRegexp.FindAllStringSubmatchIndex(n.nameFormat, -1)
	for _, m := range matches {
		fullStart, fullEnd := m[0], m[1]
		labelStart, labelEnd := m[2], m[3]

		if fullStart > cursor {
			segments = append(segments, n.nameFormat[cursor:fullStart])
		}

		key := n.nameFormat[labelStart:labelEnd]
		var value string
		switch key {
		case "#namespace":
			value = n.namespace
		case "#subsystem":
			value = n.subsystem
		case "#name":
			value = n.name
		case "qname":
			value = n.FullyQualifiedName()
		default:
			var ok bool
			value, ok = labels[key]
			if !ok {
				panic(fmt.Sprintf("invalid label in name format: %s", key))
			}
			// 将 . | : \s 全部替换为 _
			value = invalidLabelValueRegexp.ReplaceAllString(value, "_")
		}
		segments = append(segments, value)
		cursor = fullEnd
	}

	if cursor != len(n.nameFormat) {
		segments = append(segments, n.nameFormat[cursor:])
	}

	return strings.Join(segments, "")
}

func sliceToSet(set []string) map[string]struct{} {
	labelSet := make(map[string]struct{})
	for _, s := range set {
		labelSet[s] = struct{}{}
	}
	return labelSet
}
