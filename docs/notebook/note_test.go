package notebook

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexp(t *testing.T) {
	str := "Windows98 WindowsXP Windows7 Windows11"

	capture := regexp.MustCompile("Windows(XP|7|10)")
	captureRes := capture.FindAllStringSubmatchIndex(str, -1)
	require.Equal(t, 2, len(captureRes))
	require.Equal(t, 4, len(captureRes[0]))
	require.Equal(t, "WindowsXP", str[captureRes[0][0]:captureRes[0][1]])
	require.Equal(t, "XP", str[captureRes[0][2]:captureRes[0][3]])
	require.Equal(t, 4, len(captureRes[1]))
	require.Equal(t, "Windows7", str[captureRes[1][0]:captureRes[1][1]])
	require.Equal(t, "7", str[captureRes[1][2]:captureRes[1][3]])
	t.Log(captureRes)

	nonCapture := regexp.MustCompile("Windows(?:XP|7|10)")
	nonCaptureRes := nonCapture.FindAllStringSubmatchIndex(str, -1)
	require.Equal(t, 2, len(nonCaptureRes))
	require.Equal(t, 2, len(nonCaptureRes[0]))
	require.Equal(t, "WindowsXP", str[nonCaptureRes[0][0]:nonCaptureRes[0][1]])
	require.Equal(t, 2, len(nonCaptureRes[1]))
	require.Equal(t, "Windows7", str[nonCaptureRes[1][0]:nonCaptureRes[1][1]])
	t.Log(nonCaptureRes)
}

func TestRegexp2(t *testing.T) {
	r := regexp.MustCompile(`%{(color|id|level|message|module|shortfunc|time)(?::(.*?))?}`)

	res := r.FindAllStringSubmatchIndex("%{color:d}%{message:s}", -1)
	t.Log(res)
}
