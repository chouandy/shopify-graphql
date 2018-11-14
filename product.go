package shopifygraphql

// ProductCreateInput product create input struct
type ProductCreateInput struct {
	CollectionsToJoin      []string              `json:"collectionsToJoin,omitempty"`
	CollectionsToLeave     []string              `json:"collectionsToLeave,omitempty"`
	DescriptionHTML        string                `json:"descriptionHtml,omitempty"`
	GiftCardTemplateSuffix string                `json:"giftCardTemplateSuffix,omitempty"`
	Handle                 string                `json:"handle,omitempty"`
	ID                     string                `json:"id,omitempty"`
	Images                 []ImageInput          `json:"images,omitempty"`
	Metafields             []MetafieldInput      `json:"metafields,omitempty"`
	Options                []string              `json:"options,omitempty"`
	ProductType            string                `json:"productType,omitempty"`
	RedirectNewHandle      bool                  `json:"redirectNewHandle,omitempty"`
	SEO                    *SEOInput             `json:"seo,omitempty"`
	Tags                   []string              `json:"tags,omitempty"`
	TemplateSuffix         string                `json:"templateSuffix,omitempty"`
	Title                  string                `json:"title,omitempty"`
	Variants               []ProductVariantInput `json:"variants,omitempty"`
	Vendor                 string                `json:"vendor,omitempty"`
}
