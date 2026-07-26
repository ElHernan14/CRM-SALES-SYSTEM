<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';

import {
  ArrowRight,
  Building2,
  Grid2X2,
  ImageIcon,
  PackageCheck,
  Search,
  ShieldCheck,
  ShoppingBag,
  Sparkles,
  Store,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { useCategories } from '@/modules/categories/composables/useCategories';

import { useStoreProducts } from '@/modules/marketplace/composables/useStoreProducts';

import { getProductImageUrl } from '@/shared/utils/assets';

const router = useRouter();

const productParams = computed(() => ({
  page: 1,
  limit: 6,

  sort_column: 'created_at' as const,
  order: 'desc' as const,
}));

const { data: productsData, isLoading: productsLoading } = useStoreProducts(productParams);

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const featuredProducts = computed(() => {
  return productsData.value?.items ?? [];
});

const productCategories = computed(() => {
  return categoriesData.value?.product_categories?.slice(0, 6) ?? [];
});

const showcaseProducts = computed(() => {
  return featuredProducts.value.slice(0, 3);
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    maximumFractionDigits: 0,
  }).format(value);
}

function openCatalogWithCategory(categoryId: number) {
  router.push({
    name: 'store-catalog',
    query: {
      category_id: String(categoryId),
    },
  });
}

function openProduct(productId: number) {
  router.push({
    name: 'store-product',
    params: {
      productId,
    },
  });
}
</script>

<template>
  <div>
    <!-- HERO -->
    <section class="relative overflow-hidden border-b border-border bg-background">
      <div
        class="pointer-events-none absolute -left-40 top-16 h-[420px] w-[420px] rounded-full bg-primary/10 blur-3xl"
      />

      <div
        class="pointer-events-none absolute -right-32 -top-32 h-[480px] w-[480px] rounded-full bg-blue-500/10 blur-3xl"
      />

      <div
        class="relative mx-auto grid min-h-[660px] w-full max-w-[1500px] items-center gap-14 px-6 py-16 lg:grid-cols-[1.05fr_0.95fr] lg:px-8"
      >
        <!-- HERO COPY -->
        <div>
          <div
            class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/5 px-3 py-1.5 text-xs font-semibold text-primary"
          >
            <Store class="h-3.5 w-3.5" />
            Nexora commerce network
          </div>

          <h1
            class="mt-7 max-w-3xl text-5xl font-semibold tracking-[-0.055em] text-foreground sm:text-6xl lg:text-7xl"
          >
            Discover what connected businesses
            <span class="text-muted-foreground"> have to offer. </span>
          </h1>

          <p class="mt-6 max-w-2xl text-lg leading-8 text-muted-foreground">
            Browse products and services from real companies operating throughout the Nexora
            ecosystem.
          </p>

          <!-- PRIMARY SEARCH -->
          <button
            type="button"
            class="group mt-9 flex w-full max-w-2xl items-center gap-4 rounded-[1.25rem] border border-border bg-card p-2 pl-5 text-left shadow-lg transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-xl"
            @click="
              router.push({
                name: 'store-catalog',
              })
            "
          >
            <Search
              class="h-5 w-5 shrink-0 text-muted-foreground transition group-hover:text-primary"
            />

            <span class="min-w-0 flex-1 truncate text-sm text-muted-foreground">
              Search products, categories or businesses
            </span>

            <span
              class="rounded-xl bg-primary px-5 py-3 text-sm font-semibold text-primary-foreground"
            >
              Search
            </span>
          </button>

          <div class="mt-5 flex flex-col gap-3 sm:flex-row">
            <Button
              size="lg"
              class="rounded-full px-7"
              @click="
                router.push({
                  name: 'store-catalog',
                })
              "
            >
              Explore catalog
              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>

            <Button
              size="lg"
              variant="outline"
              class="rounded-full px-7"
              @click="
                router.push({
                  name: 'store-catalog',
                  query: {
                    focus: 'categories',
                  },
                })
              "
            >
              <Grid2X2 class="mr-2 h-4 w-4" />
              Browse categories
            </Button>
          </div>

          <!-- TRUST LINE -->
          <div class="mt-10 grid max-w-2xl gap-4 border-t border-border pt-6 sm:grid-cols-3">
            <div class="flex items-center gap-3">
              <ShieldCheck class="h-5 w-5 text-primary" />

              <div>
                <p class="text-sm font-semibold">Connected sellers</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Real Nexora businesses</p>
              </div>
            </div>

            <div class="flex items-center gap-3">
              <PackageCheck class="h-5 w-5 text-primary" />

              <div>
                <p class="text-sm font-semibold">Live inventory</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Availability from ERP</p>
              </div>
            </div>

            <div class="flex items-center gap-3">
              <ShoppingBag class="h-5 w-5 text-primary" />

              <div>
                <p class="text-sm font-semibold">Unified checkout</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Multi-seller purchases</p>
              </div>
            </div>
          </div>
        </div>

        <!-- PRODUCT SHOWCASE -->
        <div class="relative min-h-[480px]">
          <div class="absolute inset-8 rounded-[2.5rem] bg-primary/10 blur-3xl" />

          <div v-if="productsLoading" class="relative grid h-full min-h-[480px] grid-cols-2 gap-4">
            <div class="row-span-2 animate-pulse rounded-[2rem] bg-muted" />

            <div class="animate-pulse rounded-[2rem] bg-muted" />

            <div class="animate-pulse rounded-[2rem] bg-muted" />
          </div>

          <div
            v-else-if="showcaseProducts.length > 0"
            class="relative grid min-h-[480px] grid-cols-2 gap-4"
          >
            <button
              v-for="(product, index) in showcaseProducts"
              :key="product.id"
              type="button"
              class="group relative overflow-hidden rounded-[2rem] border border-border bg-card text-left shadow-xl transition hover:-translate-y-1 hover:shadow-2xl"
              :class="index === 0 ? 'row-span-2' : ''"
              @click="openProduct(product.id)"
            >
              <img
                v-if="getProductImageUrl(product.image_path)"
                :src="getProductImageUrl(product.image_path)!"
                :alt="product.name"
                class="absolute inset-0 h-full w-full object-cover transition duration-700 group-hover:scale-105"
              />

              <div
                v-else
                class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-muted to-background"
              >
                <ImageIcon class="h-10 w-10 text-muted-foreground/40" />
              </div>

              <div
                class="absolute inset-0 bg-gradient-to-t from-background via-background/15 to-transparent"
              />

              <div class="absolute inset-x-0 bottom-0 p-5">
                <span
                  class="inline-flex rounded-full bg-background/85 px-2.5 py-1 text-[11px] font-semibold text-foreground backdrop-blur"
                >
                  {{ product.category }}
                </span>

                <p class="mt-3 line-clamp-2 text-lg font-semibold tracking-tight">
                  {{ product.name }}
                </p>

                <div class="mt-2 flex items-end justify-between gap-3">
                  <p class="truncate text-xs text-muted-foreground">
                    {{ product.company_name }}
                  </p>

                  <p class="shrink-0 text-sm font-semibold">
                    {{ formatCurrency(product.price) }}
                  </p>
                </div>
              </div>
            </button>
          </div>

          <div
            v-else
            class="relative flex min-h-[480px] items-center justify-center rounded-[2rem] border border-border bg-card p-10 text-center shadow-xl"
          >
            <div>
              <Sparkles class="mx-auto h-8 w-8 text-primary" />

              <p class="mt-5 text-lg font-semibold">The Nexora catalog is growing</p>

              <p class="mt-2 max-w-sm text-sm leading-6 text-muted-foreground">
                Products published by connected businesses will appear here.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CATEGORY DISCOVERY -->
    <section id="store-categories" class="mx-auto w-full max-w-[1500px] px-6 py-20 lg:px-8">
      <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Browse by category
          </p>

          <h2 class="mt-3 text-3xl font-semibold tracking-[-0.04em]">
            Find the right product faster.
          </h2>

          <p class="mt-3 max-w-2xl text-sm leading-6 text-muted-foreground">
            Explore the product categories available throughout the Nexora commerce network.
          </p>
        </div>

        <Button
          variant="ghost"
          class="self-start rounded-full sm:self-auto"
          @click="
            router.push({
              name: 'store-catalog',
            })
          "
        >
          View full catalog
          <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
      </div>

      <div v-if="categoriesLoading" class="mt-9 grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <div v-for="index in 6" :key="index" class="h-36 animate-pulse rounded-[1.5rem] bg-muted" />
      </div>

      <div v-else class="mt-9 grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <button
          v-for="(category, index) in productCategories"
          :key="category.id"
          type="button"
          class="group relative overflow-hidden rounded-[1.5rem] border border-border bg-card p-5 text-left shadow-sm transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-lg"
          @click="openCatalogWithCategory(category.id)"
        >
          <div
            class="absolute right-0 top-0 h-28 w-28 rounded-full bg-primary/5 blur-2xl transition group-hover:bg-primary/10"
          />

          <div class="relative">
            <div
              class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10 text-primary"
            >
              <Building2 v-if="index % 3 === 0" class="h-5 w-5" />

              <Grid2X2 v-else-if="index % 3 === 1" class="h-5 w-5" />

              <Sparkles v-else class="h-5 w-5" />
            </div>

            <div class="mt-6 flex items-start justify-between gap-4">
              <div>
                <p class="text-base font-semibold">
                  {{ category.name }}
                </p>

                <p class="mt-2 line-clamp-2 text-sm leading-6 text-muted-foreground">
                  {{ category.description || 'Explore products from this category.' }}
                </p>
              </div>

              <ArrowRight
                class="mt-1 h-4 w-4 shrink-0 text-muted-foreground transition group-hover:translate-x-1 group-hover:text-primary"
              />
            </div>
          </div>
        </button>
      </div>
    </section>

    <section class="mx-auto w-full max-w-[1500px] px-6 pb-20 lg:px-8">
      <div
        class="relative overflow-hidden rounded-[2rem] border border-border bg-card p-7 shadow-sm sm:p-9"
      >
        <div
          class="pointer-events-none absolute -right-20 -top-24 h-64 w-64 rounded-full bg-primary/10 blur-3xl"
        />

        <div class="relative flex flex-col gap-6 lg:flex-row lg:items-center lg:justify-between">
          <div class="max-w-2xl">
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Business discovery
            </p>

            <h2 class="mt-3 text-3xl font-semibold tracking-[-0.04em]">
              Meet the businesses behind the catalog.
            </h2>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              Explore companies by industry, understand their identity and browse products directly
              from each seller.
            </p>
          </div>

          <Button
            size="lg"
            class="self-start rounded-full lg:self-auto"
            @click="
              router.push({
                name: 'store-businesses',
              })
            "
          >
            Explore businesses

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </div>
      </div>
    </section>

    <!-- FEATURED PRODUCTS -->
    <section class="border-y border-border bg-muted/15">
      <div class="mx-auto w-full max-w-[1500px] px-6 py-20 lg:px-8">
        <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Connected inventory
            </p>

            <h2 class="mt-3 text-3xl font-semibold tracking-[-0.04em]">
              Recently available on Nexora.
            </h2>

            <p class="mt-3 max-w-2xl text-sm leading-6 text-muted-foreground">
              Products and services published directly from active business workspaces.
            </p>
          </div>

          <Button
            variant="outline"
            class="self-start rounded-full sm:self-auto"
            @click="
              router.push({
                name: 'store-catalog',
              })
            "
          >
            Explore everything
            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </div>

        <div v-if="productsLoading" class="mt-9 grid gap-5 sm:grid-cols-2 xl:grid-cols-3">
          <div
            v-for="index in 6"
            :key="index"
            class="h-[420px] animate-pulse rounded-[1.75rem] bg-muted"
          />
        </div>

        <div v-else class="mt-9 grid gap-5 sm:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="product in featuredProducts"
            :key="product.id"
            class="group overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm transition hover:-translate-y-1 hover:border-primary/25 hover:shadow-xl"
          >
            <button type="button" class="block w-full text-left" @click="openProduct(product.id)">
              <div class="relative h-64 overflow-hidden bg-muted/30">
                <img
                  v-if="getProductImageUrl(product.image_path)"
                  :src="getProductImageUrl(product.image_path)!"
                  :alt="product.name"
                  class="h-full w-full object-cover transition duration-700 group-hover:scale-105"
                />

                <div
                  v-else
                  class="flex h-full items-center justify-center bg-gradient-to-br from-muted to-background"
                >
                  <ImageIcon class="h-10 w-10 text-muted-foreground/40" />
                </div>

                <div
                  class="absolute inset-0 bg-gradient-to-t from-background/70 via-transparent to-transparent"
                />

                <span
                  class="absolute left-4 top-4 rounded-full border border-white/10 bg-background/85 px-3 py-1 text-xs font-semibold backdrop-blur"
                >
                  {{ product.category }}
                </span>
              </div>

              <div class="p-5">
                <p
                  class="flex items-center gap-2 truncate text-xs font-medium text-muted-foreground"
                >
                  <Building2 class="h-3.5 w-3.5" />
                  {{ product.company_name }}
                </p>

                <h3 class="mt-3 line-clamp-2 min-h-14 text-lg font-semibold tracking-tight">
                  {{ product.name }}
                </h3>

                <div class="mt-5 flex items-end justify-between gap-4">
                  <div>
                    <p class="text-xs text-muted-foreground">Price</p>

                    <p class="mt-1 text-xl font-semibold tracking-tight">
                      {{ formatCurrency(product.price) }}
                    </p>
                  </div>

                  <div class="text-right">
                    <p class="text-xs text-muted-foreground">Availability</p>

                    <p
                      class="mt-1 text-sm font-semibold"
                      :class="
                        product.available_stock > 0
                          ? 'text-emerald-600 dark:text-emerald-400'
                          : 'text-destructive'
                      "
                    >
                      {{
                        product.available_stock > 0
                          ? `${product.available_stock} available`
                          : 'Unavailable'
                      }}
                    </p>
                  </div>
                </div>
              </div>
            </button>
          </article>
        </div>
      </div>
    </section>
  </div>
</template>
