# Shopify GraphQL

The simple GraphQL client for Shopify.

For more information, see package [`github.com/shurcooL/graphql`](https://github.com/shurcooL/graphql), which provides a GraphQL client implementation.

## Installation

```bash
go get -u github.com/chouandy/shopify-graphql
```

## Usage

Construct a Shopify GraphQL client, specifying the Shopify shop domain and storefront access token. Then, you can use it to make GraphQL queries and mutations.

```Go
// Admin API Client
client := shopifygraphql.NewAdminClient("example.myshopify.com", "admin_access_token")
// Storefront API Client
client := shopifygraphql.NewStorefrontClient("example.myshopify.com", "storefront_access_token")

// Use client...
```
