package gildedrose_test

//
// import (
//
//	"testing"
//
//	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
//
// )
//
//	func Test_Foo(t *testing.T) {
//		var items = []*gildedrose.Item{
//			{"foo", 0, 0},
//		}
//
//		gildedrose.UpdateQuality(items)
//
//		if items[0].Name != "fixme" {
//			t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
//		}
//	}

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gr "github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	t.Run("Default item", func(t *testing.T) {
		item := &gr.Item{"Foo", 10, 20}
		expected := &gr.Item{"Foo", 9, 19}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Default item with SellIn passed and Quality above zero", func(t *testing.T) {
		item := &gr.Item{"Bar", 0, 20}
		expected := &gr.Item{"Bar", -1, 18}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Default item with Quality zero", func(t *testing.T) {
		item := &gr.Item{"Baz", 10, 0}
		expected := &gr.Item{"Baz", 9, 0}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Aged Brie", func(t *testing.T) {
		item := &gr.Item{"Aged Brie", 10, 20}
		expected := &gr.Item{"Aged Brie", 9, 21}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Aged Brie with Quality 50", func(t *testing.T) {
		item := &gr.Item{"Aged Brie", 10, 50}
		expected := &gr.Item{"Aged Brie", 9, 50}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Aged Brie with SellIn passed", func(t *testing.T) {
		item := &gr.Item{"Aged Brie", 0, 20}
		expected := &gr.Item{"Aged Brie", -1, 22}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Backstage passes", func(t *testing.T) {
		item := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 20, 20}
		expected := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 19, 21}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Backstage passes with Quality 50", func(t *testing.T) {
		item := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 20, 50}
		expected := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 19, 50}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Backstage passes with low SellIn", func(t *testing.T) {
		item := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20}
		expected := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 9, 22}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Backstage passes with very low SellIn", func(t *testing.T) {
		item := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 5, 20}
		expected := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 4, 23}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Backstage passes with SellIn passed", func(t *testing.T) {
		item := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", 0, 20}
		expected := &gr.Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})

	t.Run("Sulfuras", func(t *testing.T) {
		item := &gr.Item{"Sulfuras, Hand of Ragnaros", 10, 80}
		expected := &gr.Item{"Sulfuras, Hand of Ragnaros", 10, 80}
		gr.UpdateQuality([]*gr.Item{item})
		assert.Equal(t, expected, item)
	})
}
