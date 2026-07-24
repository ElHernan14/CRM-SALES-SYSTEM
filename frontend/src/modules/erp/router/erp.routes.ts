import ERPLayout from '@/modules/erp/layouts/ERPLayout.vue';
import DashboardPage from '@/modules/dashboard/pages/DashboardPage.vue';

const ProductsPage = () => import('@/modules/products/pages/ProductsPage.vue');
const InvoicesPage = () => import('@/modules/invoices/pages/InvoicesPage.vue');
const PaymentsPage = () => import('@/modules/erp/pages/PaymentsPage.vue');
const MarketplacePage = () => import('@/modules/marketplace/pages/MarketplacePage.vue');
const SupplierCatalogPage = () => import('@/modules/marketplace/pages/SupplierCatalogPage.vue');
const PurchasesPage = () => import('@/modules/purchases/pages/PurchasesPage.vue');
const InventoryPage = () => import('@/modules/erp/pages/InventoryPage.vue');
const CustomersPage = () => import('@/modules/erp/pages/CustomersPage.vue');
const AnalyticsPage = () => import('@/modules/erp/pages/AnalyticsPage.vue');
const SettingsPage = () => import('@/modules/company/pages/CompanySettingsPage.vue');

export const erpRoutes = [
  {
    path: '/erp',
    component: ERPLayout,
    meta: {
      requiresAuth: true,
      requiresCompany: true,
    },
    children: [
      { path: 'dashboard', name: 'erp-dashboard', component: DashboardPage },
      { path: 'products', name: 'erp-products', component: ProductsPage },
      { path: 'sales', name: 'erp-sales', component: InvoicesPage },
      { path: 'payments', name: 'erp-payments', component: PaymentsPage },
      {
        path: 'settings',
        component: () => import('@/modules/settings/pages/ERPSettingsPage.vue'),
        meta: {
          requiresAuth: true,
        },
        children: [
          {
            path: '',
            redirect: {
              name: 'erp-settings-company',
            },
          },
          {
            path: 'company',
            name: 'erp-settings-company',
            component: SettingsPage,
            props: {
              embedded: true,
            },
          },
          {
            path: 'profile',
            name: 'erp-settings-profile',
            component: () => import('@/modules/account/pages/StoreAccountPage.vue'),
            props: {
              embedded: true,
              surface: 'erp',
            },
          },
        ],
      },
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
