package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullyQualifiedName(t *testing.T) {
	n1 := &Namer{
		namespace:  "namespace",
		subsystem:  "subsystem",
		name:       "name",
	}
	require.Equal(t, "namespace.subsystem.name", n1.FullyQualifiedName())

	n2 := &Namer{
		subsystem:  "subsystem",
		name:       "name",
	}
	require.Equal(t, "subsystem.name", n2.FullyQualifiedName())

	n3 := &Namer{
		namespace:  "namespace",
		name:       "name",
	}
	require.Equal(t, "namespace.name", n3.FullyQualifiedName())

	n4 := &Namer{
		name:       "name",
	}
	require.Equal(t, "name", n4.FullyQualifiedName())
}

func TestFormat(t *testing.T) {
	labels := []string{"name", "age", "school", "home"}
	n := &Namer{
		namespace:  "namespace",
		subsystem:  "subsystem",
		name:       "name",
		nameFormat: "%{#namespace}.%{#subsystem}_%{#name}..%{school}",
		labelNames: sliceToSet(labels),
	}

	f := n.Format("name", "xiangyu wu", "age", "18", "school", "78-senior high school", "home", "yaohai")
	t.Log(f)
}