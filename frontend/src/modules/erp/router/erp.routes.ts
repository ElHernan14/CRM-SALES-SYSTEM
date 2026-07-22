import ERPLayout from '@/modules/erp/layouts/ERPLayout.vue';
import DashboardPage from '@/modules/erp/pages/DashboardPage.vue';

const ProductsPage = () => import('@/modules/products/pages/ProductsPage.vue');
const InvoicesPage = () => import('@/modules/invoices/pages/InvoicesPage.vue');
const PaymentsPage = () => import('@/modules/erp/pages/PaymentsPage.vue');
const MarketplacePage = () => import('@/modules/marketplace/pages/MarketplacePage.vue');
const SupplierCatalogPage = () => import('@/modules/marketplace/pages/SupplierCatalogPage.vue');
const PurchasesPage = () => import('@/modules/purchases/pages/PurchasesPage.vue');
const InventoryPage = () => import('@/modules/erp/pages/InventoryPage.vue');
const CustomersPage = () => import('@/modules/erp/pages/CustomersPage.vue');
const AnalyticsPage = () => import('@/modules/erp/pages/AnalyticsPage.vue');
const SettingsPage = () => import('@/modules/erp/pages/SettingsPage.vue');

export const erpRoutes = [
  {
    path: '/erp',
    component: ERPLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      { path: 'dashboard', name: 'erp-dashboard', component: DashboardPage },
      { path: 'products', name: 'erp-products', component: ProductsPage },
      { path: 'invoices', name: 'erp-invoices', component: InvoicesPage },
      { path: 'payments', name: 'erp-payments', component: PaymentsPage },
      { path: 'settings', name: 'erp-settings', component: SettingsPage },
      { path: 'marketplace', name: 'erp-marketplace', component: MarketplacePage },
      {
        path: 'marketplace/suppliers/:supplierId',
        name: 'erp-supplier-catalog',
        component: SupplierCatalogPage,
      },
      { path: 'purchases', name: 'erp-purchases', component: PurchasesPage },
      { path: 'inventory', name: 'erp-inventory', component: InventoryPage },
      { path: 'customers', name: 'erp-customers', component: CustomersPage },
      { path: 'analytics', name: 'erp-analytics', component: AnalyticsPage },
    ],
  },
];
