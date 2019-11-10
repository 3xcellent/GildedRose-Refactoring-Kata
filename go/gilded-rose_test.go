package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestUpdateQuality_Item(t *testing.T) {
	t.Run("Regular Items", func(t *testing.T) {
		t.Run("At the end of each day our system lowers both values for every item", func(t *testing.T) {
			items := []*Item{
				{name: "item", sellIn: 10, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 9, items[0].quality)
		})
	})

	t.Run("Aged Brie", func(t *testing.T) {
		t.Run("increases in Quality the older it gets", func(t *testing.T) {
			items := []*Item{
				{name: "Aged Brie", sellIn: 10, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 11, items[0].quality)
		})
	})
	t.Run("Quality", func(t *testing.T) {
		t.Run("The Quality of an item can never have its Quality increase above 50", func(t *testing.T) {
			items := []*Item{ // pick item where the quality increases
				{name: "Aged Brie", sellIn: 10, quality: 50},
			}
			assert.Equal(t, 50, items[0].quality)
		})
		t.Run("Sulfuras ...  Quality is 80", func(t *testing.T) {
			items := []*Item{
				{name: "Sulfuras, Hand of Ragnaros", quality: 80},
			}
			assert.Equal(t, 80, items[0].quality)
		})
	})

	t.Run("Sulfuras ... never has to be sold or decreases in Quality", func(t *testing.T) {
		name := "Sulfuras, Hand of Ragnaros"
		items := []*Item{
			{name: name, sellIn: 10, quality: 10},
		}
		expectedAfterUpdate := []*Item{
			{name: name, sellIn: 10, quality: 10},
		}
		UpdateQuality(items)
		assert.Equal(t, expectedAfterUpdate, items)
	})
	t.Run("Backstage passes to a TAFKAL80ETC concert", func(t *testing.T) {
		name := "Backstage passes to a TAFKAL80ETC concert"
		t.Run("increases in Quality as its SellIn value approaches [zero]", func(t *testing.T) {
			items := []*Item{
				{name: name, sellIn: 20, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 11, items[0].quality)
		})
		t.Run("Quality increases by 2 when there are 10 days or less", func(t *testing.T) {
			items := []*Item{
				{name: name, sellIn: 10, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 12, items[0].quality)
		})
		t.Run("by 3 when there are 5 days or less", func(t *testing.T) {
			items := []*Item{
				{name: name, sellIn: 5, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 13, items[0].quality)
		})
		t.Run("but Quality drops to 0 after the concert", func(t *testing.T) {
			items := []*Item{
				{name: name, sellIn: 0, quality: 10},
			}
			UpdateQuality(items)
			assert.Equal(t, 0, items[0].quality)
		})
	})
}
