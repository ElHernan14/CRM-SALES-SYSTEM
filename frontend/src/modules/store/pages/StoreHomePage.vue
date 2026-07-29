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

function openCatalog() {
  router.push({
    name: 'store-catalog',
  });
}

function openBusinesses() {
  router.push({
    name: 'store-businesses',
  });
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
        class="mx-auto grid min-h-[560px] w-full max-w-[1500px] items-center gap-10 px-4 py-14 sm:px-6 sm:py-16 lg:min-h-[620px] lg:grid-cols-[1.1fr_0.9fr] lg:gap-12 lg:px-8"
      >
        <!-- HERO COPY -->
        <div class="min-w-0">
          <div
            class="inline-flex max-w-full items-center gap-2 rounded-full border border-primary/20 bg-primary/5 px-3 py-1.5 text-xs font-semibold text-primary"
          >
            <Store class="h-3.5 w-3.5 shrink-0" />

            <span class="truncate"> Nexora commerce network </span>
          </div>

          <h1
            class="mt-5 max-w-3xl text-[2.65rem] font-semibold leading-[1.03] tracking-[-0.045em] text-foreground sm:mt-6 sm:text-6xl lg:text-7xl"
          >
            Discover what connected businesses

            <span class="text-muted-foreground"> have to offer. </span>
          </h1>

          <p
            class="mt-5 max-w-2xl text-base leading-7 text-muted-foreground sm:mt-6 sm:text-lg sm:leading-8"
          >
            Browse products and services from real companies operating throughout the Nexora
            ecosystem.
          </p>

          <!-- PRIMARY SEARCH -->
          <button
            type="button"
            class="group mt-8 flex w-full max-w-2xl items-center gap-3 rounded-[1.25rem] border border-border bg-card p-2 pl-4 text-left shadow-lg transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-xl sm:mt-9 sm:gap-4 sm:pl-5"
            @click="openCatalog"
          >
            <Search
              class="h-5 w-5 shrink-0 text-muted-foreground transition group-hover:text-primary"
            />

            <span class="min-w-0 flex-1 truncate text-sm text-muted-foreground">
              Search products, categories or businesses
            </span>

            <span
              class="shrink-0 rounded-xl bg-primary px-4 py-3 text-sm font-semibold text-primary-foreground sm:px-5"
            >
              <span class="hidden min-[390px]:inline"> Search </span>

              <ArrowRight class="h-4 w-4 min-[390px]:hidden" />
            </span>
          </button>

          <div class="mt-8 flex flex-col gap-3 sm:flex-row">
            <Button size="lg" class="w-full rounded-full px-7 sm:w-auto" @click="openCatalog">
              Explore catalog

              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>

            <Button
              size="lg"
              variant="outline"
              class="w-full rounded-full px-7 sm:w-auto"
              @click="openBusinesses"
            >
              <Building2 class="mr-2 h-4 w-4" />

              Browse businesses
            </Button>
          </div>

          <!-- TRUST LINE -->
          <div
            class="mt-9 grid gap-4 text-sm text-muted-foreground sm:grid-cols-3 sm:gap-5 lg:flex lg:flex-wrap lg:gap-x-8 lg:gap-y-4"
          >
            <div class="flex min-w-0 items-center gap-3">
              <ShieldCheck class="h-5 w-5 shrink-0 text-primary" />

              <div class="min-w-0">
                <p class="text-sm font-semibold text-foreground">Connected sellers</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Real Nexora businesses</p>
              </div>
            </div>

            <div class="flex min-w-0 items-center gap-3">
              <PackageCheck class="h-5 w-5 shrink-0 text-primary" />

              <div class="min-w-0">
                <p class="text-sm font-semibold text-foreground">Live inventory</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Availability from ERP</p>
              </div>
            </div>

            <div class="flex min-w-0 items-center gap-3">
              <ShoppingBag class="h-5 w-5 shrink-0 text-primary" />

              <div class="min-w-0">
                <p class="text-sm font-semibold text-foreground">Unified checkout</p>

                <p class="mt-0.5 text-xs text-muted-foreground">Multi-seller purchases</p>
              </div>
            </div>
          </div>
        </div>

        <!-- PRODUCT SHOWCASE -->
        <div
          class="relative min-h-[340px] overflow-hidden rounded-[1.5rem] border border-border bg-card p-3 shadow-2xl sm:min-h-[420px] sm:rounded-[2rem] sm:p-4"
        >
          <div
            class="pointer-events-none absolute inset-8 rounded-[2.5rem] bg-primary/10 blur-3xl"
          />

          <!-- SHOWCASE LOADING -->
          <div
            v-if="productsLoading"
            class="relative min-h-[314px] sm:grid sm:min-h-[448px] sm:grid-cols-2 sm:gap-4"
          >
            <div
              class="h-[314px] animate-pulse rounded-[1.5rem] bg-muted sm:row-span-2 sm:h-auto sm:rounded-[2rem]"
            />

            <div class="hidden animate-pulse rounded-[2rem] bg-muted sm:block" />

            <div class="hidden animate-pulse rounded-[2rem] bg-muted sm:block" />
          </div>

          <!-- SHOWCASE PRODUCTS -->
          <div
            v-else-if="showcaseProducts.length > 0"
            class="relative grid min-h-[314px] w-full grid-cols-1 sm:min-h-[448px] sm:grid-cols-2 sm:gap-4"
          >
            <button
              v-for="(product, index) in showcaseProducts"
              :key="product.id"
              type="button"
              class="group relative block w-full min-w-0 overflow-hidden rounded-[1.5rem] border border-border bg-card text-left shadow-xl transition hover:-translate-y-1 hover:shadow-2xl sm:rounded-[2rem]"
              :class="[index === 0 ? 'h-[314px] sm:row-span-2 sm:h-auto' : 'hidden sm:block']"
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
                class="absolute inset-0 bg-gradient-to-t from-background via-background/20 to-transparent"
              />

              <div class="absolute inset-x-0 bottom-0 min-w-0 p-4 sm:p-5">
                <span
                  class="inline-flex max-w-full rounded-full bg-background/85 px-2.5 py-1 text-[11px] font-semibold text-foreground backdrop-blur"
                >
                  <span class="truncate">
                    {{ product.category }}
                  </span>
                </span>

                <p class="mt-3 line-clamp-2 text-lg font-semibold tracking-tight">
                  {{ product.name }}
                </p>

                <div class="mt-2 flex min-w-0 items-end justify-between gap-3">
                  <p class="min-w-0 truncate text-xs text-muted-foreground">
                    {{ product.company_name }}
                  </p>

                  <p class="shrink-0 text-sm font-semibold">
                    {{ formatCurrency(product.price) }}
                  </p>
                </div>
              </div>
            </button>
          </div>

          <!-- SHOWCASE EMPTY -->
          <div
            v-else
            class="relative flex min-h-[314px] items-center justify-center rounded-[1.5rem] border border-border bg-card p-6 text-center shadow-xl sm:min-h-[448px] sm:rounded-[2rem] sm:p-10"
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
    <section
      id="store-categories"
      class="mx-auto w-full max-w-[1500px] px-4 py-16 sm:px-6 sm:py-20 lg:px-8"
    >
      <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Browse by category
          </p>

          <h2 class="mt-3 text-3xl font-semibold tracking-[-0.04em] sm:text-4xl">
            Find the right product faster.
          </h2>

          <p class="mt-3 max-w-2xl text-sm leading-6 text-muted-foreground">
            Explore the product categories available throughout the Nexora commerce network.
          </p>
        </div>

        <Button
          variant="ghost"
          class="w-full self-start rounded-full sm:w-auto sm:self-auto"
          @click="openCatalog"
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
          class="group relative min-w-0 overflow-hidden rounded-[1.5rem] border border-border bg-card p-5 text-left shadow-sm transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-lg"
          @click="openCatalogWithCategory(category.id)"
        >
          <div
            class="pointer-events-none absolute right-0 top-0 h-28 w-28 rounded-full bg-primary/5 blur-2xl transition group-hover:bg-primary/10"
          />

          <div class="relative">
            <div
              class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10 text-primary"
            >
              <Building2 v-if="index % 3 === 0" class="h-5 w-5" />

              <Grid2X2 v-else-if="index % 3 === 1" class="h-5 w-5" />

              <Sparkles v-else class="h-5 w-5" />
            </div>

            <div class="mt-6 flex min-w-0 items-start justify-between gap-4">
              <div class="min-w-0">
                <p class="break-words text-base font-semibold">
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

    <!-- BUSINESS DISCOVERY -->
    <section class="mx-auto w-full max-w-[1500px] px-4 pb-16 sm:px-6 sm:pb-20 lg:px-8">
      <div
        class="nexora-surface nexora-glow relative grid overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-xl lg:grid-cols-[0.92fr_1.08fr]"
      >
        <!-- COPY -->
        <div class="relative z-10 flex flex-col justify-center p-6 sm:p-9 lg:p-12">
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Business discovery
          </p>

          <h2 class="mt-3 max-w-xl text-3xl font-semibold tracking-[-0.04em] sm:text-4xl">
            Meet the businesses behind the catalog.
          </h2>

          <p class="mt-4 max-w-xl text-sm leading-6 text-muted-foreground">
            Explore companies by industry, understand their identity and browse products directly
            from each connected seller.
          </p>

          <div class="mt-7">
            <Button size="lg" class="w-full rounded-full sm:w-auto" @click="openBusinesses">
              Explore businesses

              <ArrowRight
                class="ml-2 h-4 w-4 transition-transform duration-200 group-hover:translate-x-0.5"
              />
            </Button>
          </div>

          <div class="mt-8 grid gap-4 border-t border-border pt-6 min-[420px]:grid-cols-2">
            <div class="flex items-start gap-3">
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Building2 class="h-4 w-4 text-primary" />
              </div>

              <div>
                <p class="text-sm font-semibold">Verified businesses</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Real companies connected to Nexora.
                </p>
              </div>
            </div>

            <div class="flex items-start gap-3">
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <PackageCheck class="h-4 w-4 text-primary" />
              </div>

              <div>
                <p class="text-sm font-semibold">Connected inventory</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Catalog availability backed by ERP data.
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- IMAGE -->
        <div class="relative min-h-[280px] overflow-hidden sm:min-h-[360px] lg:min-h-[440px]">
          <img
            src="/images/store/store-discovery.jpg"
            alt="Modern warehouse operations connected to the Nexora commerce network"
            class="absolute inset-0 h-full w-full object-cover transition duration-700 hover:scale-[1.025]"
            loading="lazy"
            decoding="async"
          />

          <!-- Mobile readability -->
          <div
            class="absolute inset-0 bg-gradient-to-t from-card/75 via-transparent to-transparent lg:hidden"
          />

          <!-- Desktop connection with copy -->
          <div
            class="absolute inset-0 hidden bg-gradient-to-r from-card via-card/20 to-transparent lg:block"
          />

          <!-- Nexora visual identity -->
          <div
            class="absolute inset-0 bg-gradient-to-br from-violet-950/10 via-transparent to-primary/15"
          />

          <div
            class="pointer-events-none absolute -right-24 -top-24 h-72 w-72 rounded-full bg-violet-500/20 blur-3xl"
          />

          <div class="absolute inset-0 ring-1 ring-inset ring-white/10" />

          <div
            class="absolute bottom-5 left-5 right-5 flex items-center justify-between rounded-2xl border border-white/10 bg-slate-950/55 px-4 py-3 text-white shadow-xl backdrop-blur-md sm:bottom-6 sm:left-6 sm:right-6"
          >
            <div>
              <p class="text-xs text-slate-300">Connected operations</p>

              <p class="mt-1 text-sm font-semibold">Inventory visible throughout Nexora</p>
            </div>

            <Store class="h-5 w-5 shrink-0 text-blue-300" />
          </div>
        </div>
      </div>
    </section>

    <!-- FEATURED PRODUCTS -->
    <section class="border-y border-border bg-muted/15">
      <div class="mx-auto w-full max-w-[1500px] px-4 py-16 sm:px-6 sm:py-20 lg:px-8">
        <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Connected inventory
            </p>

            <h2 class="mt-3 text-3xl font-semibold tracking-[-0.04em] sm:text-4xl">
              Recently available on Nexora.
            </h2>

            <p class="mt-3 max-w-2xl text-sm leading-6 text-muted-foreground">
              Products and services published directly from active business workspaces.
            </p>
          </div>

          <Button
            variant="outline"
            class="w-full self-start rounded-full sm:w-auto sm:self-auto"
            @click="openCatalog"
          >
            Explore everything

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </div>

        <div v-if="productsLoading" class="mt-9 grid gap-5 sm:grid-cols-2 xl:grid-cols-3">
          <div
            v-for="index in 6"
            :key="index"
            class="h-[390px] animate-pulse rounded-[1.75rem] bg-muted sm:h-[420px]"
          />
        </div>

        <div v-else class="mt-9 grid gap-5 sm:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="product in featuredProducts"
            :key="product.id"
            class="group min-w-0 overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm transition hover:-translate-y-1 hover:border-primary/25 hover:shadow-xl"
          >
            <button
              type="button"
              class="block w-full min-w-0 text-left"
              @click="openProduct(product.id)"
            >
              <div class="relative h-52 overflow-hidden bg-muted/30 sm:h-64">
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
                  class="absolute left-4 top-4 max-w-[calc(100%-2rem)] truncate rounded-full border border-white/10 bg-background/85 px-3 py-1 text-xs font-semibold backdrop-blur"
                >
                  {{ product.category }}
                </span>
              </div>

              <div class="min-w-0 p-5">
                <p
                  class="flex min-w-0 items-center gap-2 text-xs font-medium text-muted-foreground"
                >
                  <Building2 class="h-3.5 w-3.5 shrink-0" />

                  <span class="truncate">
                    {{ product.company_name }}
                  </span>
                </p>

                <h3 class="mt-3 line-clamp-2 min-h-14 text-lg font-semibold tracking-tight">
                  {{ product.name }}
                </h3>

                <div
                  class="mt-5 flex flex-col gap-4 border-t border-border pt-4 min-[390px]:flex-row min-[390px]:items-end min-[390px]:justify-between"
                >
                  <div class="min-w-0">
                    <p class="text-xs text-muted-foreground">Price</p>

                    <p class="mt-1 truncate text-xl font-semibold tracking-tight">
                      {{ formatCurrency(product.price) }}
                    </p>
                  </div>

                  <div class="min-w-0 min-[390px]:text-right">
                    <p class="text-xs text-muted-foreground">Availability</p>

                    <p
                      class="mt-1 truncate text-sm font-semibold"
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
