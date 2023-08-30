package mansort_test

import (
	"dev03/mansort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []string{
	"1 дек 6 14:29 Android",
	"2 июн 10 2015 Sources",
	"3 окт 31 15:08 VirtualBox",
	"4 янв 13 11:42 Lightworks",
	"5 янв 11 12:33 Pictures",
}

func TestSort_EmptyInput(t *testing.T) {
	lines := mansort.Sort([]string{}, -1, false, false, false)
	if len(lines) != 0 {
		t.Errorf("expected 0 lines, got %d", len(lines))
	}
}

func TestSort_EmptyParams(t *testing.T) {
	lines := mansort.Sort(testData, -1, false, false, false)
	assert.Equal(t, []string{
		"1 дек 6 14:29 Android",
		"2 июн 10 2015 Sources",
		"3 окт 31 15:08 VirtualBox",
		"4 янв 13 11:42 Lightworks",
		"5 янв 11 12:33 Pictures",
	}, lines)
}

func TestSort_WithColumn(t *testing.T) {
	lines := mansort.Sort(testData, 5, false, false, false)
	assert.Equal(t, []string{
		"1 дек 6 14:29 Android",
		"4 янв 13 11:42 Lightworks",
		"5 янв 11 12:33 Pictures",
		"2 июн 10 2015 Sources",
		"3 окт 31 15:08 VirtualBox",
	}, lines)
}

func TestSort_WithNumericColumn(t *testing.T) {
	lines := mansort.Sort(testData, 3, true, false, false)
	assert.Equal(t, []string{
		"1 дек 6 14:29 Android",
		"2 июн 10 2015 Sources",
		"5 янв 11 12:33 Pictures",
		"4 янв 13 11:42 Lightworks",
		"3 окт 31 15:08 VirtualBox",
	}, lines)
}

func TestSort_WithReverse(t *testing.T) {
	lines := mansort.Sort(testData, 3, true, true, false)
	assert.Equal(t, []string{
		"3 окт 31 15:08 VirtualBox",
		"4 янв 13 11:42 Lightworks",
		"5 янв 11 12:33 Pictures",
		"2 июн 10 2015 Sources",
		"1 дек 6 14:29 Android",
	}, lines)
}

func TestSort_WithUnique(t *testing.T) {
	lines := mansort.Sort([]string{
		"4",
		"1",
		"2",
		"1",
		"8",
	}, -1, true, false, true)
	assert.Equal(t, []string{
		"1",
		"2",
		"4",
		"8",
	}, lines)
}
