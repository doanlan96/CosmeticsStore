package crawler

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

func CrawlProduct(db *sql.DB)  {
	c := colly.NewCollector()
	c.OnHTML(".product", func(e *colly.HTMLElement) {
		product := models.Product{}
		product.ProductName = e.ChildText(".card-title a")
		priceTypeTex := e.ChildAttr(".event_buy_now_add_to_cart", "data-product-price")
		product.Price, _ = strconv.ParseFloat(priceTypeTex, 64)
		product.Image = e.ChildAttr(".card-image", "src")
		product.Description = "Capacity:\n\nSK-II Facial Treatment Cleanser 120g: deep cleansing, brightening, bactericidal\n\n- Facial Treatment Clear Lotion 160ml: clean the skin, balance the skin before moisturizing\n\n SK-II Facial Treatment Essence 75ml: moisturizing, brightening, tightening pores, creating the best environment for serums and creams to penetrate\n\n - Anti-aging essence SK-II RNA Radical New Age Essence 50ml: anti-aging, remove wrinkles, tighten pores\n\n - SK-II Cellumination Deep Surge Ex 50g whitening cream: moisturizing, brightening pink, blurring dark spots, smoothing skin, tightening pores\n\n \n\nCommitment:\n\n- Genuine product SK-II Japan\n\n- The product when delivered to you has enough stamps and boxes\n\n- Product has a shelf life of at least 20 months when delivered\n\n \n\nINSTANT 6% OFF AND GET GENUINE SK-II JAPANESE BAG VALUED 200K WHEN BUYING A HUGE SIZE IMMEDIATE SKIN CLEANSER!\n\n \n\n\n\n \n\nLarge deep whitening set to help whiten skin naturally, fade dark spots, prevent the appearance of melanin and tighten pores include:\n\n \n\n1. SK-II SKII FACIAL TREATMENT GENTLE CLEANSER 120G face wash View More Details\n\n\n \n\nSK-II facial cleanser with ingredients containing Pitera and white willow extract helps to completely remove impurities.\nCreamy texture for soft and hydrated face.\nIt can also remove makeup completely.\nAfter just one use, your skin will look noticeably brighter.\n \n\nUser manual\n\nUse twice a day morning and night\nApply warm water to your face, take about 1 cm of Facial Treatment Gentle Cleanser in the palm of your hand and add a small amount of water. Whip the cleanser until foam forms. Spread the foam evenly on the face and pat the face from bottom to top. Rinse with warm water.\n \n\n2. SK-II FACIAL TREATMENT CLEAR LOTION 160ML View more details\n\n\nSK-II toner is very safe for all skin types, many women inbox to ask the shop if acne or sensitive skin can be used, don't worry. Moreover, SK-II toner also contains Salicylic Acid, an oil-soluble compound that penetrates deep into the pores to clear all excess, helping to unclog pores.\n\nMost of us think that rose water is only for cleaning purposes, but we do not know that it is extremely important in skin care. Let's take a look at the great effects that SK-II rose water brings.\n\npH balance\nDetoxification and Purification\nShrink and tighten pores\nReplenish water and nutrition for the skin\n    \n\nUser manual\n\nUse twice a day morning and night\nSoak a cotton pad with the solution and apply all over the face, especially the oily T-zone. Continue to apply to the areas from the chin to the bottom of the chin and cheeks. Then apply to the neck area in an upward direction to avoid sagging skin. Use after washing your face, before using magic water."
		product.IsHot = false
		product.CategoryName = "Face Scream"
		product.CategoryName = "SK-II"
		product.CategoryId = 1
		product.BrandId = 1
		product.NumberAvailable = 100
		fmt.Printf("- Name: %s- Price: %f- Picture: %s", product.ProductName, product.Price, product.Image) //In ra màn hình giá trị đã lấy được
		stmt, err := database.DB.Prepare("INSERT INTO products(product_name, description, price, image, is_hot, category_id, brand_id, number_available) VALUES (?,?,?,?,?,?,?,?)")
		fmt.Println(err)
		res, err := stmt.Exec(product.ProductName, product.Description, product.Price, product.Image, product.IsHot, product.CategoryId, product.BrandId, product.NumberAvailable)
		fmt.Println(err)

		lastId, err := res.LastInsertId()
		//Lấy ra ID vừa được insert

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("=&gt;Insert ID: %dnn", lastId)
	})
	c.OnScraped(func(r *colly.Response) { //Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.sk-ii.com/our-products")
}