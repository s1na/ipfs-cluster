// Code generated by "stringer -type=phase"; DO NOT EDIT.

package maptracker

import "strconv"

const _Phase_name = "PhaseErrorPhaseQueuedPhaseInProgress"

var _Phase_index = [...]uint8{0, 10, 21, 36}

func (i phase) String() string {
	if i < 0 || i >= phase(len(_Phase_index)-1) {
		return "phase(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Phase_name[_Phase_index[i]:_Phase_index[i+1]]
}