<script setup lang="ts">
import LandingHeader from '../components/LandingHeader.vue';
import LandingHero from '../components/LandingHero.vue';

import OperationalFoundation from '../components/OperationalFoundation.vue';
import NexoraLoop from '../components/NexoraLoop.vue';

import BusinessWorkspaceSection from '../components/BusinessWorkspaceSection.vue';
import ConnectedCommerceSection from '../components/ConnectedCommerceSection.vue';

import CommerceModesSection from '../components/CommerceModesSection.vue';
import InvoiceLifecycleSection from '../components/InvoiceLifecycleSection.vue';

import ArchitectureSection from '../components/ArchitectureSection.vue';
import ProductExperienceSection from '../components/ProductExperienceSection.vue';

import LandingCta from '../components/LandingCta.vue';
import LandingFooter from '../components/LandingFooter.vue';
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <LandingHeader />

    <main>
      <LandingHero />

      <OperationalFoundation />

      <NexoraLoop />

      <BusinessWorkspaceSection />

      <ConnectedCommerceSection />

      <CommerceModesSection />

      <InvoiceLifecycleSection />

      <ArchitectureSection />

      <ProductExperienceSection />

      <LandingCta />
    </main>

    <LandingFooter />
  </div>
</template>
