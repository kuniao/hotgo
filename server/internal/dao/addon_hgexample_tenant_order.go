// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAddonHgexampleTenantOrderDao is internal type for wrapping internal DAO implements.
type internalAddonHgexampleTenantOrderDao = *internal.AddonHgexampleTenantOrderDao

// addonHgexampleTenantOrderDao is the data access object for table hg_addon_hgexample_tenant_order.
// You can define custom methods on it to extend its functionality as you wish.
type addonHgexampleTenantOrderDao struct {
	internalAddonHgexampleTenantOrderDao
}

var (
	// AddonHgexampleTenantOrder is globally public accessible object for table hg_addon_hgexample_tenant_order operations.
	AddonHgexampleTenantOrder = addonHgexampleTenantOrderDao{
		internal.NewAddonHgexampleTenantOrderDao(),
	}
)

// Fill with you ideas below.