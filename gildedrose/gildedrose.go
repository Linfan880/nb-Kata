package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

//
//	func UpdateQuality(items []*Item) {
//		for i := 0; i < len(items); i++ {
//
//			if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
//				if items[i].Quality > 0 {
//					if items[i].Name != "Sulfuras, Hand of Ragnaros" {
//						items[i].Quality = items[i].Quality - 1
//					}
//				}
//			} else {
//				if items[i].Quality < 50 {
//					items[i].Quality = items[i].Quality + 1
//					if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
//						if items[i].SellIn < 11 {
//							if items[i].Quality < 50 {
//								items[i].Quality = items[i].Quality + 1
//							}
//						}
//						if items[i].SellIn < 6 {
//							if items[i].Quality < 50 {
//								items[i].Quality = items[i].Quality + 1
//							}
//						}
//					}
//				}
//			}
//
//			if items[i].Name != "Sulfuras, Hand of Ragnaros" {
//				items[i].SellIn = items[i].SellIn - 1
//			}
//
//			if items[i].SellIn < 0 {
//				if items[i].Name != "Aged Brie" {
//					if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
//						if items[i].Quality > 0 {
//							if items[i].Name != "Sulfuras, Hand of Ragnaros" {
//								items[i].Quality = items[i].Quality - 1
//							}
//						}
//					} else {
//						items[i].Quality = items[i].Quality - items[i].Quality
//					}
//				} else {
//					if items[i].Quality < 50 {
//						items[i].Quality = items[i].Quality + 1
//					}
//				}
//			}
//		}
//
// }

type ItemUpdater interface {
	Update()
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		itemUpdater := getItemUpdater(item)
		itemUpdater.Update()
	}
}

func getItemUpdater(item *Item) ItemUpdater {
	switch item.Name {
	case "Aged Brie":
		return NewAgedBrie(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePasses(item)
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(item)
	default:
		return NewDefaultItem(item)
	}
}

type DefaultItem struct {
	item *Item
}

func NewDefaultItem(item *Item) *DefaultItem {
	return &DefaultItem{item: item}
}

func (item *DefaultItem) Update() {
	if item.item.Quality > 0 {
		item.item.Quality--
	}

	item.item.SellIn--

	if item.item.SellIn < 0 && item.item.Quality > 0 {
		item.item.Quality--
	}
}

type AgedBrie struct {
	item *Item
}

func NewAgedBrie(item *Item) *AgedBrie {
	return &AgedBrie{item: item}
}

func (item *AgedBrie) Update() {
	if item.item.Quality < 50 {
		item.item.Quality++
	}

	item.item.SellIn--

	if item.item.SellIn < 0 && item.item.Quality < 50 {
		item.item.Quality++
	}
}

type BackstagePasses struct {
	item *Item
}

func NewBackstagePasses(item *Item) *BackstagePasses {
	return &BackstagePasses{item: item}
}

func (item *BackstagePasses) Update() {
	if item.item.Quality < 50 {
		item.item.Quality++

		if item.item.SellIn < 11 {
			item.item.Quality++
		}
		if item.item.SellIn < 6 {
			item.item.Quality++
		}
	}

	item.item.SellIn--

	if item.item.SellIn < 0 {
		item.item.Quality = 0
	}
}

type Sulfuras struct {
	item *Item
}

func NewSulfuras(item *Item) *Sulfuras {
	return &Sulfuras{item: item}
}

func (item *Sulfuras) Update() {}
