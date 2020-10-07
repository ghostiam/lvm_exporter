package collector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// --reportformat json -o all --units B --binary
func TestReportsToMetrics(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []MetricsLabels
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				b: []byte(`{"report": []}`),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "lv all fields",
			args: args{
				b: []byte(`{
      "report": [
          {
              "lv": [
                  {"lv_uuid":"rKAgQS-1WI6-H43d-amB7-IGFr-kJOI-KhNAuA", "lv_name":"lvmirror", "lv_full_name":"vgdata/lvmirror", "lv_path":"/dev/vgdata/lvmirror", "lv_dm_path":"/dev/mapper/vgdata-lvmirror", "lv_parent":"", "lv_layout":"raid,raid1", "lv_role":"public", "lv_initial_image_sync":"1", "lv_image_synced":"0", "lv_merging":"0", "lv_converting":"0", "lv_allocation_policy":"inherit", "lv_allocation_locked":"0", "lv_fixed_minor":"0", "lv_skip_activation":"0", "lv_when_full":"", "lv_active":"active", "lv_active_locally":"1", "lv_active_remotely":"0", "lv_active_exclusively":"1", "lv_major":"-1", "lv_minor":"-1", "lv_read_ahead":"auto", "lv_size":"3221225472B", "lv_metadata_size":"", "seg_count":"1", "origin":"", "origin_uuid":"", "origin_size":"", "lv_ancestors":"", "lv_full_ancestors":"", "lv_descendants":"", "lv_full_descendants":"", "raid_mismatch_count":"0", "raid_sync_action":"idle", "raid_write_behind":"", "raid_min_recovery_rate":"", "raid_max_recovery_rate":"", "move_pv":"", "move_pv_uuid":"", "convert_lv":"", "convert_lv_uuid":"", "mirror_log":"", "mirror_log_uuid":"", "data_lv":"", "data_lv_uuid":"", "metadata_lv":"", "metadata_lv_uuid":"", "pool_lv":"", "pool_lv_uuid":"", "lv_tags":"", "lv_profile":"", "lv_lockargs":"", "lv_time":"2020-10-05 17:24:40 +0000", "lv_time_removed":"", "lv_host":"LVM-TEST", "lv_modules":"raid", "lv_historical":"0", "lv_kernel_major":"253", "lv_kernel_minor":"4", "lv_kernel_read_ahead":"131072B", "lv_permissions":"writeable", "lv_suspended":"0", "lv_live_table":"1", "lv_inactive_table":"0", "lv_device_open":"0", "data_percent":"", "snap_percent":"", "metadata_percent":"", "copy_percent":"100.00", "sync_percent":"100.00", "cache_total_blocks":"", "cache_used_blocks":"", "cache_dirty_blocks":"", "cache_read_hits":"", "cache_read_misses":"", "cache_write_hits":"", "cache_write_misses":"", "kernel_cache_settings":"", "kernel_cache_policy":"", "kernel_metadata_format":"", "lv_health_status":"", "kernel_discards":"", "lv_check_needed":"-1", "lv_merge_failed":"-1", "lv_snapshot_invalid":"-1", "vdo_operating_mode":"", "vdo_compression_state":"", "vdo_index_state":"", "vdo_used_size":"", "vdo_saving_percent":"", "lv_attr":"rwi-a-r---", "vg_fmt":"lvm2", "vg_uuid":"GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4", "vg_name":"vgdata", "vg_attr":"wz--n-", "vg_permissions":"writeable", "vg_extendable":"1", "vg_exported":"0", "vg_partial":"0", "vg_allocation_policy":"normal", "vg_clustered":"0", "vg_shared":"0", "vg_size":"21466447872B", "vg_free":"13941866496B", "vg_sysid":"", "vg_systemid":"", "vg_lock_type":"", "vg_lock_args":"", "vg_extent_size":"4194304B", "vg_extent_count":"5118", "vg_free_count":"3324", "max_lv":"0", "max_pv":"0", "pv_count":"2", "vg_missing_pv_count":"0", "lv_count":"2", "snap_count":"0", "vg_seqno":"32", "vg_tags":"", "vg_profile":"", "vg_mda_count":"2", "vg_mda_used_count":"2", "vg_mda_free":"518656B", "vg_mda_size":"1044480B", "vg_mda_copies":"unmanaged"},
                  {"lv_uuid":"i30Cyh-D0cc-Rnll-EqHY-WawI-12P0-1tZ4XJ", "lv_name":"lvol0", "lv_full_name":"vgdata/lvol0", "lv_path":"/dev/vgdata/lvol0", "lv_dm_path":"/dev/mapper/vgdata-lvol0", "lv_parent":"", "lv_layout":"linear", "lv_role":"public", "lv_initial_image_sync":"0", "lv_image_synced":"0", "lv_merging":"0", "lv_converting":"0", "lv_allocation_policy":"inherit", "lv_allocation_locked":"0", "lv_fixed_minor":"0", "lv_skip_activation":"0", "lv_when_full":"", "lv_active":"active", "lv_active_locally":"1", "lv_active_remotely":"0", "lv_active_exclusively":"1", "lv_major":"-1", "lv_minor":"-1", "lv_read_ahead":"auto", "lv_size":"1073741824B", "lv_metadata_size":"", "seg_count":"1", "origin":"", "origin_uuid":"", "origin_size":"", "lv_ancestors":"", "lv_full_ancestors":"", "lv_descendants":"", "lv_full_descendants":"", "raid_mismatch_count":"", "raid_sync_action":"", "raid_write_behind":"", "raid_min_recovery_rate":"", "raid_max_recovery_rate":"", "move_pv":"", "move_pv_uuid":"", "convert_lv":"", "convert_lv_uuid":"", "mirror_log":"", "mirror_log_uuid":"", "data_lv":"", "data_lv_uuid":"", "metadata_lv":"", "metadata_lv_uuid":"", "pool_lv":"", "pool_lv_uuid":"", "lv_tags":"", "lv_profile":"", "lv_lockargs":"", "lv_time":"2020-10-05 18:06:13 +0000", "lv_time_removed":"", "lv_host":"LVM-TEST", "lv_modules":"", "lv_historical":"0", "lv_kernel_major":"253", "lv_kernel_minor":"5", "lv_kernel_read_ahead":"131072B", "lv_permissions":"writeable", "lv_suspended":"0", "lv_live_table":"1", "lv_inactive_table":"0", "lv_device_open":"0", "data_percent":"", "snap_percent":"", "metadata_percent":"", "copy_percent":"", "sync_percent":"", "cache_total_blocks":"", "cache_used_blocks":"", "cache_dirty_blocks":"", "cache_read_hits":"", "cache_read_misses":"", "cache_write_hits":"", "cache_write_misses":"", "kernel_cache_settings":"", "kernel_cache_policy":"", "kernel_metadata_format":"", "lv_health_status":"", "kernel_discards":"", "lv_check_needed":"-1", "lv_merge_failed":"-1", "lv_snapshot_invalid":"-1", "vdo_operating_mode":"", "vdo_compression_state":"", "vdo_index_state":"", "vdo_used_size":"", "vdo_saving_percent":"", "lv_attr":"-wi-a-----", "vg_fmt":"lvm2", "vg_uuid":"GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4", "vg_name":"vgdata", "vg_attr":"wz--n-", "vg_permissions":"writeable", "vg_extendable":"1", "vg_exported":"0", "vg_partial":"0", "vg_allocation_policy":"normal", "vg_clustered":"0", "vg_shared":"0", "vg_size":"21466447872B", "vg_free":"13941866496B", "vg_sysid":"", "vg_systemid":"", "vg_lock_type":"", "vg_lock_args":"", "vg_extent_size":"4194304B", "vg_extent_count":"5118", "vg_free_count":"3324", "max_lv":"0", "max_pv":"0", "pv_count":"2", "vg_missing_pv_count":"0", "lv_count":"2", "snap_count":"0", "vg_seqno":"32", "vg_tags":"", "vg_profile":"", "vg_mda_count":"2", "vg_mda_used_count":"2", "vg_mda_free":"518656B", "vg_mda_size":"1044480B", "vg_mda_copies":"unmanaged"}
              ]
          }
      ]
  }`),
			},
			want: []MetricsLabels{
				{
					Namespace: "lv",
					Metrics: []Metric{
						{Field: "cache_dirty_blocks", Value: -1, Help: "CacheDirtyBlocks - Dirty cache blocks."},
						{Field: "cache_read_hits", Value: -1, Help: "CacheReadHits - Cache read hits."},
						{Field: "cache_read_misses", Value: -1, Help: "CacheReadMisses - Cache read misses."},
						{Field: "cache_total_blocks", Value: -1, Help: "CacheTotalBlocks - Total cache blocks."},
						{Field: "cache_used_blocks", Value: -1, Help: "CacheUsedBlocks - Used cache blocks."},
						{Field: "cache_write_hits", Value: -1, Help: "CacheWriteHits - Cache write hits."},
						{Field: "cache_write_misses", Value: -1, Help: "CacheWriteMisses - Cache write misses."},
						{Field: "copy_percent", Value: 100, Help: "Cpy%Sync - For Cache, RAID, mirrors and pvmove, current percentage in-sync."},
						{Field: "data_percent", Value: -1, Help: "Data% - For snapshot, cache and thin pools and volumes, the percentage full if LV is active."},
						{Field: "kernel_metadata_format", Value: -1, Help: "KMFmt - Cache metadata format used in kernel."},
						{Field: "lv_active_exclusively", Value: 1, Help: "ActExcl - Set if the LV is active exclusively."},
						{Field: "lv_active_locally", Value: 1, Help: "ActLocal - Set if the LV is active locally."},
						{Field: "lv_active_remotely", Value: 0, Help: "ActRemote - Set if the LV is active remotely."},
						{Field: "lv_allocation_locked", Value: 0, Help: "AllocLock - Set if LV is locked against allocation changes."},
						{Field: "lv_check_needed", Value: -1, Help: "CheckNeeded - For thin pools and cache volumes, whether metadata check is needed."},
						{Field: "lv_converting", Value: 0, Help: "Converting - Set if LV is being converted."},
						{Field: "lv_count", Value: 2, Help: "#LV - Number of LVs."},
						{Field: "lv_device_open", Value: 0, Help: "DevOpen - Set if LV device is open."},
						{Field: "lv_fixed_minor", Value: 0, Help: "FixMin - Set if LV has fixed minor number assigned."},
						{Field: "lv_health_status", Value: 0, Help: "Health - LV health status. -1: undefined, 0: , 1: partial"},
						{Field: "lv_historical", Value: 0, Help: "Historical - Set if the LV is historical."},
						{Field: "lv_image_synced", Value: 0, Help: "ImgSynced - Set if mirror/RAID image is synchronized."},
						{Field: "lv_inactive_table", Value: 0, Help: "InactiveTable - Set if LV has inactive table present."},
						{Field: "lv_initial_image_sync", Value: 1, Help: "InitImgSync - Set if mirror/RAID images underwent initial resynchronization."},
						{Field: "lv_kernel_major", Value: 253, Help: "KMaj - Currently assigned major number or -1 if LV is not active."},
						{Field: "lv_kernel_minor", Value: 4, Help: "KMin - Currently assigned minor number or -1 if LV is not active."},
						{Field: "lv_kernel_read_ahead", Value: 131072, Help: "KRahead - Currently-in-use read ahead setting in current units."},
						{Field: "lv_live_table", Value: 1, Help: "LiveTable - Set if LV has live table present."},
						{Field: "lv_major", Value: -1, Help: "Maj - Persistent major number or -1 if not persistent."},
						{Field: "lv_merge_failed", Value: -1, Help: "MergeFailed - Set if snapshot merge failed."},
						{Field: "lv_merging", Value: 0, Help: "Merging - Set if snapshot LV is being merged to origin."},
						{Field: "lv_metadata_size", Value: -1, Help: "MSize - For thin and cache pools, the size of the LV that holds the metadata."},
						{Field: "lv_minor", Value: -1, Help: "Min - Persistent minor number or -1 if not persistent."},
						{Field: "lv_permissions", Value: 0, Help: "LPerms - LV permissions. -1: undefined, 0: writeable, 1: read-only, 2: read-only-override"},
						{Field: "lv_read_ahead", Value: -1, Help: "Rahead - Read ahead setting in current units."},
						{Field: "lv_size", Value: 3.221225472e+09, Help: "LSize - Size of LV in current units."},
						{Field: "lv_skip_activation", Value: 0, Help: "SkipAct - Set if LV is skipped on activation."},
						{Field: "lv_snapshot_invalid", Value: -1, Help: "SnapInvalid - Set if snapshot LV is invalid."},
						{Field: "lv_suspended", Value: 0, Help: "Suspended - Set if LV is suspended."},
						{Field: "lv_time", Value: 1.60191868e+09, Help: "CTime - Creation time of the LV, if known"},
						{Field: "lv_time_removed", Value: 0, Help: "RTime - Removal time of the LV, if known"},
						{Field: "lv_when_full", Value: -1, Help: "WhenFull - For thin pools, behavior when full. -1: undefined, 0: error, 1: queue"},
						{Field: "max_lv", Value: 0, Help: "MaxLV - Maximum number of LVs allowed in VG or 0 if unlimited."},
						{Field: "max_pv", Value: 0, Help: "MaxPV - Maximum number of PVs allowed in VG or 0 if unlimited."},
						{Field: "metadata_percent", Value: -1, Help: "Meta% - For cache and thin pools, the percentage of metadata full if LV is active."},
						{Field: "origin_size", Value: -1, Help: "OSize - For snapshots, the size of the origin device of this LV."},
						{Field: "pv_count", Value: 2, Help: "#PV - Number of PVs in VG."},
						{Field: "raid_max_recovery_rate", Value: -1, Help: "MaxSync - For RAID1, the maximum recovery I/O load in kiB/sec/disk."},
						{Field: "raid_min_recovery_rate", Value: -1, Help: "MinSync - For RAID1, the minimum recovery I/O load in kiB/sec/disk."},
						{Field: "raid_mismatch_count", Value: 0, Help: "Mismatches - For RAID, number of mismatches found or repaired."},
						{Field: "raid_sync_action", Value: 0, Help: "SyncAction - For RAID, the current synchronization action being performed. -1: undefined, 0: idle, 1: frozen, 2: resync, 3: recover, 4: check, 5: repair"},
						{Field: "raid_write_behind", Value: -1, Help: "WBehind - For RAID1, the number of outstanding writes allowed to writemostly devices."},
						{Field: "seg_count", Value: 1, Help: "#Seg - Number of segments in LV."},
						{Field: "snap_count", Value: 0, Help: "#SN - Number of snapshots."},
						{Field: "snap_percent", Value: -1, Help: "Snap% - For snapshots, the percentage full if LV is active."},
						{Field: "sync_percent", Value: 100, Help: "Cpy%Sync - For Cache, RAID, mirrors and pvmove, current percentage in-sync."},
						{Field: "vdo_saving_percent", Value: -1, Help: "VDOSaving% - For vdo pools, percentage of saved space."},
						{Field: "vdo_used_size", Value: -1, Help: "VDOUsedSize - For vdo pools, currently used space."},
						{Field: "vg_allocation_policy", Value: 0, Help: "AllocPol - VG allocation policy. -1: undefined, 0: normal, 1: contiguous, 2: cling, 3: anywhere, 4: inherited"},
						{Field: "vg_clustered", Value: 0, Help: "Clustered - Set if VG is clustered."},
						{Field: "vg_exported", Value: 0, Help: "Exported - Set if VG is exported."},
						{Field: "vg_extendable", Value: 1, Help: "Extendable - Set if VG is extendable."},
						{Field: "vg_extent_count", Value: 5118, Help: "#Ext - Total number of Physical Extents."},
						{Field: "vg_extent_size", Value: 4.194304e+06, Help: "Ext - Size of Physical Extents in current units."},
						{Field: "vg_free", Value: 1.3941866496e+10, Help: "VFree - Total amount of free space in current units."},
						{Field: "vg_free_count", Value: 3324, Help: "Free - Total number of unallocated Physical Extents."},
						{Field: "vg_mda_copies", Value: -1, Help: "#VMdaCps - Target number of in use metadata areas in the VG."},
						{Field: "vg_mda_count", Value: 2, Help: "#VMda - Number of metadata areas on this VG."},
						{Field: "vg_mda_free", Value: 518656, Help: "VMdaFree - Free metadata area space for this VG in current units."},
						{Field: "vg_mda_size", Value: 1.04448e+06, Help: "VMdaSize - Size of smallest metadata area for this VG in current units."},
						{Field: "vg_mda_used_count", Value: 2, Help: "#VMdaUse - Number of metadata areas in use on this VG."},
						{Field: "vg_missing_pv_count", Value: 0, Help: "#PV Missing - Number of PVs in VG which are missing."},
						{Field: "vg_partial", Value: 0, Help: "Partial - Set if VG is partial."},
						{Field: "vg_permissions", Value: 0, Help: "VPerms - VG permissions. -1: undefined, 0: writeable, 1: read-only"},
						{Field: "vg_seqno", Value: 32, Help: "Seq - Revision number of internal metadata.  Incremented whenever it changes."},
						{Field: "vg_shared", Value: 0, Help: "Shared - Set if VG is shared."},
						{Field: "vg_size", Value: 2.1466447872e+10, Help: "VSize - Total size of VG in current units."},
					},
					Labels: []Label{
						{Field: "lv_host", Value: "LVM-TEST"},
						{Field: "lv_layout", Value: "raid,raid1"},
						{Field: "lv_name", Value: "lvmirror"},
						{Field: "lv_uuid", Value: "rKAgQS-1WI6-H43d-amB7-IGFr-kJOI-KhNAuA"},
						{Field: "vg_name", Value: "vgdata"},
						{Field: "vg_uuid", Value: "GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4"},
					},
				},
				{
					Namespace: "lv",
					Metrics: []Metric{
						{Field: "cache_dirty_blocks", Value: -1, Help: "CacheDirtyBlocks - Dirty cache blocks."},
						{Field: "cache_read_hits", Value: -1, Help: "CacheReadHits - Cache read hits."},
						{Field: "cache_read_misses", Value: -1, Help: "CacheReadMisses - Cache read misses."},
						{Field: "cache_total_blocks", Value: -1, Help: "CacheTotalBlocks - Total cache blocks."},
						{Field: "cache_used_blocks", Value: -1, Help: "CacheUsedBlocks - Used cache blocks."},
						{Field: "cache_write_hits", Value: -1, Help: "CacheWriteHits - Cache write hits."},
						{Field: "cache_write_misses", Value: -1, Help: "CacheWriteMisses - Cache write misses."},
						{Field: "copy_percent", Value: -1, Help: "Cpy%Sync - For Cache, RAID, mirrors and pvmove, current percentage in-sync."},
						{Field: "data_percent", Value: -1, Help: "Data% - For snapshot, cache and thin pools and volumes, the percentage full if LV is active."},
						{Field: "kernel_metadata_format", Value: -1, Help: "KMFmt - Cache metadata format used in kernel."},
						{Field: "lv_active_exclusively", Value: 1, Help: "ActExcl - Set if the LV is active exclusively."},
						{Field: "lv_active_locally", Value: 1, Help: "ActLocal - Set if the LV is active locally."},
						{Field: "lv_active_remotely", Value: 0, Help: "ActRemote - Set if the LV is active remotely."},
						{Field: "lv_allocation_locked", Value: 0, Help: "AllocLock - Set if LV is locked against allocation changes."},
						{Field: "lv_check_needed", Value: -1, Help: "CheckNeeded - For thin pools and cache volumes, whether metadata check is needed."},
						{Field: "lv_converting", Value: 0, Help: "Converting - Set if LV is being converted."},
						{Field: "lv_count", Value: 2, Help: "#LV - Number of LVs."},
						{Field: "lv_device_open", Value: 0, Help: "DevOpen - Set if LV device is open."},
						{Field: "lv_fixed_minor", Value: 0, Help: "FixMin - Set if LV has fixed minor number assigned."},
						{Field: "lv_health_status", Value: 0, Help: "Health - LV health status. -1: undefined, 0: , 1: partial"},
						{Field: "lv_historical", Value: 0, Help: "Historical - Set if the LV is historical."},
						{Field: "lv_image_synced", Value: 0, Help: "ImgSynced - Set if mirror/RAID image is synchronized."},
						{Field: "lv_inactive_table", Value: 0, Help: "InactiveTable - Set if LV has inactive table present."},
						{Field: "lv_initial_image_sync", Value: 0, Help: "InitImgSync - Set if mirror/RAID images underwent initial resynchronization."},
						{Field: "lv_kernel_major", Value: 253, Help: "KMaj - Currently assigned major number or -1 if LV is not active."},
						{Field: "lv_kernel_minor", Value: 5, Help: "KMin - Currently assigned minor number or -1 if LV is not active."},
						{Field: "lv_kernel_read_ahead", Value: 131072, Help: "KRahead - Currently-in-use read ahead setting in current units."},
						{Field: "lv_live_table", Value: 1, Help: "LiveTable - Set if LV has live table present."},
						{Field: "lv_major", Value: -1, Help: "Maj - Persistent major number or -1 if not persistent."},
						{Field: "lv_merge_failed", Value: -1, Help: "MergeFailed - Set if snapshot merge failed."},
						{Field: "lv_merging", Value: 0, Help: "Merging - Set if snapshot LV is being merged to origin."},
						{Field: "lv_metadata_size", Value: -1, Help: "MSize - For thin and cache pools, the size of the LV that holds the metadata."},
						{Field: "lv_minor", Value: -1, Help: "Min - Persistent minor number or -1 if not persistent."},
						{Field: "lv_permissions", Value: 0, Help: "LPerms - LV permissions. -1: undefined, 0: writeable, 1: read-only, 2: read-only-override"},
						{Field: "lv_read_ahead", Value: -1, Help: "Rahead - Read ahead setting in current units."},
						{Field: "lv_size", Value: 1.073741824e+09, Help: "LSize - Size of LV in current units."},
						{Field: "lv_skip_activation", Value: 0, Help: "SkipAct - Set if LV is skipped on activation."},
						{Field: "lv_snapshot_invalid", Value: -1, Help: "SnapInvalid - Set if snapshot LV is invalid."},
						{Field: "lv_suspended", Value: 0, Help: "Suspended - Set if LV is suspended."},
						{Field: "lv_time", Value: 1.601921173e+09, Help: "CTime - Creation time of the LV, if known"},
						{Field: "lv_time_removed", Value: 0, Help: "RTime - Removal time of the LV, if known"},
						{Field: "lv_when_full", Value: -1, Help: "WhenFull - For thin pools, behavior when full. -1: undefined, 0: error, 1: queue"},
						{Field: "max_lv", Value: 0, Help: "MaxLV - Maximum number of LVs allowed in VG or 0 if unlimited."},
						{Field: "max_pv", Value: 0, Help: "MaxPV - Maximum number of PVs allowed in VG or 0 if unlimited."},
						{Field: "metadata_percent", Value: -1, Help: "Meta% - For cache and thin pools, the percentage of metadata full if LV is active."},
						{Field: "origin_size", Value: -1, Help: "OSize - For snapshots, the size of the origin device of this LV."},
						{Field: "pv_count", Value: 2, Help: "#PV - Number of PVs in VG."},
						{Field: "raid_max_recovery_rate", Value: -1, Help: "MaxSync - For RAID1, the maximum recovery I/O load in kiB/sec/disk."},
						{Field: "raid_min_recovery_rate", Value: -1, Help: "MinSync - For RAID1, the minimum recovery I/O load in kiB/sec/disk."},
						{Field: "raid_mismatch_count", Value: -1, Help: "Mismatches - For RAID, number of mismatches found or repaired."},
						{Field: "raid_sync_action", Value: -1, Help: "SyncAction - For RAID, the current synchronization action being performed. -1: undefined, 0: idle, 1: frozen, 2: resync, 3: recover, 4: check, 5: repair"},
						{Field: "raid_write_behind", Value: -1, Help: "WBehind - For RAID1, the number of outstanding writes allowed to writemostly devices."},
						{Field: "seg_count", Value: 1, Help: "#Seg - Number of segments in LV."},
						{Field: "snap_count", Value: 0, Help: "#SN - Number of snapshots."},
						{Field: "snap_percent", Value: -1, Help: "Snap% - For snapshots, the percentage full if LV is active."},
						{Field: "sync_percent", Value: -1, Help: "Cpy%Sync - For Cache, RAID, mirrors and pvmove, current percentage in-sync."},
						{Field: "vdo_saving_percent", Value: -1, Help: "VDOSaving% - For vdo pools, percentage of saved space."},
						{Field: "vdo_used_size", Value: -1, Help: "VDOUsedSize - For vdo pools, currently used space."},
						{Field: "vg_allocation_policy", Value: 0, Help: "AllocPol - VG allocation policy. -1: undefined, 0: normal, 1: contiguous, 2: cling, 3: anywhere, 4: inherited"},
						{Field: "vg_clustered", Value: 0, Help: "Clustered - Set if VG is clustered."},
						{Field: "vg_exported", Value: 0, Help: "Exported - Set if VG is exported."},
						{Field: "vg_extendable", Value: 1, Help: "Extendable - Set if VG is extendable."},
						{Field: "vg_extent_count", Value: 5118, Help: "#Ext - Total number of Physical Extents."},
						{Field: "vg_extent_size", Value: 4.194304e+06, Help: "Ext - Size of Physical Extents in current units."},
						{Field: "vg_free", Value: 1.3941866496e+10, Help: "VFree - Total amount of free space in current units."},
						{Field: "vg_free_count", Value: 3324, Help: "Free - Total number of unallocated Physical Extents."},
						{Field: "vg_mda_copies", Value: -1, Help: "#VMdaCps - Target number of in use metadata areas in the VG."},
						{Field: "vg_mda_count", Value: 2, Help: "#VMda - Number of metadata areas on this VG."},
						{Field: "vg_mda_free", Value: 518656, Help: "VMdaFree - Free metadata area space for this VG in current units."},
						{Field: "vg_mda_size", Value: 1.04448e+06, Help: "VMdaSize - Size of smallest metadata area for this VG in current units."},
						{Field: "vg_mda_used_count", Value: 2, Help: "#VMdaUse - Number of metadata areas in use on this VG."},
						{Field: "vg_missing_pv_count", Value: 0, Help: "#PV Missing - Number of PVs in VG which are missing."},
						{Field: "vg_partial", Value: 0, Help: "Partial - Set if VG is partial."},
						{Field: "vg_permissions", Value: 0, Help: "VPerms - VG permissions. -1: undefined, 0: writeable, 1: read-only"},
						{Field: "vg_seqno", Value: 32, Help: "Seq - Revision number of internal metadata.  Incremented whenever it changes."},
						{Field: "vg_shared", Value: 0, Help: "Shared - Set if VG is shared."},
						{Field: "vg_size", Value: 2.1466447872e+10, Help: "VSize - Total size of VG in current units."},
					},
					Labels: []Label{
						{Field: "lv_host", Value: "LVM-TEST"},
						{Field: "lv_layout", Value: "linear"},
						{Field: "lv_name", Value: "lvol0"},
						{Field: "lv_uuid", Value: "i30Cyh-D0cc-Rnll-EqHY-WawI-12P0-1tZ4XJ"},
						{Field: "vg_name", Value: "vgdata"},
						{Field: "vg_uuid", Value: "GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "vg all fields",
			args: args{
				b: []byte(`{
      "report": [
          {
              "vg": [
                  {"vg_fmt":"lvm2", "vg_uuid":"GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4", "vg_name":"vgdata", "vg_attr":"wz--n-", "vg_permissions":"writeable", "vg_extendable":"1", "vg_exported":"0", "vg_partial":"0", "vg_allocation_policy":"normal", "vg_clustered":"0", "vg_shared":"0", "vg_size":"21466447872B", "vg_free":"13941866496B", "vg_sysid":"", "vg_systemid":"", "vg_lock_type":"", "vg_lock_args":"", "vg_extent_size":"4194304B", "vg_extent_count":"5118", "vg_free_count":"3324", "max_lv":"0", "max_pv":"0", "pv_count":"2", "vg_missing_pv_count":"0", "lv_count":"2", "snap_count":"0", "vg_seqno":"32", "vg_tags":"", "vg_profile":"", "vg_mda_count":"2", "vg_mda_used_count":"2", "vg_mda_free":"518656B", "vg_mda_size":"1044480B", "vg_mda_copies":"unmanaged"}
              ]
          }
      ]
  }`),
			},
			want: []MetricsLabels{
				{
					Namespace: "vg",
					Metrics: []Metric{
						{Field: "lv_count", Value: 2, Help: "#LV - Number of LVs."},
						{Field: "max_lv", Value: 0, Help: "MaxLV - Maximum number of LVs allowed in VG or 0 if unlimited."},
						{Field: "max_pv", Value: 0, Help: "MaxPV - Maximum number of PVs allowed in VG or 0 if unlimited."},
						{Field: "pv_count", Value: 2, Help: "#PV - Number of PVs in VG."},
						{Field: "snap_count", Value: 0, Help: "#SN - Number of snapshots."},
						{Field: "vg_allocation_policy", Value: 0, Help: "AllocPol - VG allocation policy. -1: undefined, 0: normal, 1: contiguous, 2: cling, 3: anywhere, 4: inherited"},
						{Field: "vg_clustered", Value: 0, Help: "Clustered - Set if VG is clustered."},
						{Field: "vg_exported", Value: 0, Help: "Exported - Set if VG is exported."},
						{Field: "vg_extendable", Value: 1, Help: "Extendable - Set if VG is extendable."},
						{Field: "vg_extent_count", Value: 5118, Help: "#Ext - Total number of Physical Extents."},
						{Field: "vg_extent_size", Value: 4194304, Help: "Ext - Size of Physical Extents in current units."},
						{Field: "vg_free", Value: 13941866496, Help: "VFree - Total amount of free space in current units."},
						{Field: "vg_free_count", Value: 3324, Help: "Free - Total number of unallocated Physical Extents."},
						{Field: "vg_mda_copies", Value: -1, Help: "#VMdaCps - Target number of in use metadata areas in the VG."},
						{Field: "vg_mda_count", Value: 2, Help: "#VMda - Number of metadata areas on this VG."},
						{Field: "vg_mda_free", Value: 518656, Help: "VMdaFree - Free metadata area space for this VG in current units."},
						{Field: "vg_mda_size", Value: 1044480, Help: "VMdaSize - Size of smallest metadata area for this VG in current units."},
						{Field: "vg_mda_used_count", Value: 2, Help: "#VMdaUse - Number of metadata areas in use on this VG."},
						{Field: "vg_missing_pv_count", Value: 0, Help: "#PV Missing - Number of PVs in VG which are missing."},
						{Field: "vg_partial", Value: 0, Help: "Partial - Set if VG is partial."},
						{Field: "vg_permissions", Value: 0, Help: "VPerms - VG permissions. -1: undefined, 0: writeable, 1: read-only"},
						{Field: "vg_seqno", Value: 32, Help: "Seq - Revision number of internal metadata.  Incremented whenever it changes."},
						{Field: "vg_shared", Value: 0, Help: "Shared - Set if VG is shared."},
						{Field: "vg_size", Value: 21466447872, Help: "VSize - Total size of VG in current units."},
					},
					Labels: []Label{
						{Field: "vg_name", Value: "vgdata"},
						{Field: "vg_uuid", Value: "GlVa5G-kmsv-G0Hu-fN1W-2p3B-Lh2f-7MY7K4"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pv all fields",
			args: args{
				b: []byte(`{
      "report": [
          {
              "pv": [
                  {"pv_fmt":"lvm2", "pv_uuid":"3SiVuY-idTg-118r-p3cc-hOYp-0qKE-6lS0Fq", "dev_size":"10737418240B", "pv_name":"/dev/sdb", "pv_major":"8", "pv_minor":"16", "pv_mda_free":"518656B", "pv_mda_size":"1044480B", "pv_ext_vsn":"2", "pe_start":"1048576B", "pv_size":"10733223936B", "pv_free":"7507804160B", "pv_used":"3225419776B", "pv_attr":"a--", "pv_allocatable":"1", "pv_exported":"0", "pv_missing":"0", "pv_pe_count":"2559", "pv_pe_alloc_count":"769", "pv_tags":"", "pv_mda_count":"1", "pv_mda_used_count":"1", "pv_ba_start":"0B", "pv_ba_size":"0B", "pv_in_use":"1", "pv_duplicate":"0"},
                  {"pv_fmt":"lvm2", "pv_uuid":"PTjDUr-PO0k-kqFA-99Nc-qqsM-bBe0-tHxm5q", "dev_size":"10737418240B", "pv_name":"/dev/sdc", "pv_major":"8", "pv_minor":"32", "pv_mda_free":"518656B", "pv_mda_size":"1044480B", "pv_ext_vsn":"2", "pe_start":"1048576B", "pv_size":"10733223936B", "pv_free":"6434062336B", "pv_used":"3225419776B", "pv_attr":"a--", "pv_allocatable":"1", "pv_exported":"0", "pv_missing":"0", "pv_pe_count":"2559", "pv_pe_alloc_count":"1025", "pv_tags":"", "pv_mda_count":"1", "pv_mda_used_count":"1", "pv_ba_start":"0B", "pv_ba_size":"0B", "pv_in_use":"1", "pv_duplicate":"0"}
              ]
          }
      ]
  }`),
			},
			want: []MetricsLabels{
				{
					Namespace: "pv",
					Metrics: []Metric{
						{Field: "dev_size", Value: 10737418240, Help: "DevSize - Size of underlying device in current units."},
						{Field: "pe_start", Value: 1048576, Help: "1st PE - Offset to the start of data on the underlying device."},
						{Field: "pv_allocatable", Value: 1, Help: "Allocatable - Set if this device can be used for allocation."},
						{Field: "pv_ba_size", Value: 0, Help: "BA Size - Size of PV Bootloader Area in current units."},
						{Field: "pv_ba_start", Value: 0, Help: "BA Start - Offset to the start of PV Bootloader Area on the underlying device in current units."},
						{Field: "pv_duplicate", Value: 0, Help: "Duplicate - Set if PV is an unchosen duplicate."},
						{Field: "pv_exported", Value: 0, Help: "Exported - Set if this device is exported."},
						{Field: "pv_ext_vsn", Value: 2, Help: "PExtVsn - PV header extension version."},
						{Field: "pv_free", Value: 7507804160, Help: "PFree - Total amount of unallocated space in current units."},
						{Field: "pv_in_use", Value: 1, Help: "PInUse - Set if PV is used."},
						{Field: "pv_mda_count", Value: 1, Help: "#PMda - Number of metadata areas on this device."},
						{Field: "pv_mda_free", Value: 518656, Help: "PMdaFree - Free metadata area space on this device in current units."},
						{Field: "pv_mda_size", Value: 1044480, Help: "PMdaSize - Size of smallest metadata area on this device in current units."},
						{Field: "pv_mda_used_count", Value: 1, Help: "#PMdaUse - Number of metadata areas in use on this device."},
						{Field: "pv_missing", Value: 0, Help: "Missing - Set if this device is missing in system."},
						{Field: "pv_pe_alloc_count", Value: 769, Help: "Alloc - Total number of allocated Physical Extents."},
						{Field: "pv_pe_count", Value: 2559, Help: "PE - Total number of Physical Extents."},
						{Field: "pv_size", Value: 10733223936, Help: "PSize - Size of PV in current units."},
						{Field: "pv_used", Value: 3225419776, Help: "Used - Total amount of allocated space in current units."},
					},
					Labels: []Label{
						{Field: "pv_name", Value: "/dev/sdb"},
						{Field: "pv_uuid", Value: "3SiVuY-idTg-118r-p3cc-hOYp-0qKE-6lS0Fq"},
					},
				},
				{
					Namespace: "pv",
					Metrics: []Metric{
						{Field: "dev_size", Value: 10737418240, Help: "DevSize - Size of underlying device in current units."},
						{Field: "pe_start", Value: 1048576, Help: "1st PE - Offset to the start of data on the underlying device."},
						{Field: "pv_allocatable", Value: 1, Help: "Allocatable - Set if this device can be used for allocation."},
						{Field: "pv_ba_size", Value: 0, Help: "BA Size - Size of PV Bootloader Area in current units."},
						{Field: "pv_ba_start", Value: 0, Help: "BA Start - Offset to the start of PV Bootloader Area on the underlying device in current units."},
						{Field: "pv_duplicate", Value: 0, Help: "Duplicate - Set if PV is an unchosen duplicate."},
						{Field: "pv_exported", Value: 0, Help: "Exported - Set if this device is exported."},
						{Field: "pv_ext_vsn", Value: 2, Help: "PExtVsn - PV header extension version."},
						{Field: "pv_free", Value: 6434062336, Help: "PFree - Total amount of unallocated space in current units."},
						{Field: "pv_in_use", Value: 1, Help: "PInUse - Set if PV is used."},
						{Field: "pv_mda_count", Value: 1, Help: "#PMda - Number of metadata areas on this device."},
						{Field: "pv_mda_free", Value: 518656, Help: "PMdaFree - Free metadata area space on this device in current units."},
						{Field: "pv_mda_size", Value: 1044480, Help: "PMdaSize - Size of smallest metadata area on this device in current units."},
						{Field: "pv_mda_used_count", Value: 1, Help: "#PMdaUse - Number of metadata areas in use on this device."},
						{Field: "pv_missing", Value: 0, Help: "Missing - Set if this device is missing in system."},
						{Field: "pv_pe_alloc_count", Value: 1025, Help: "Alloc - Total number of allocated Physical Extents."},
						{Field: "pv_pe_count", Value: 2559, Help: "PE - Total number of Physical Extents."},
						{Field: "pv_size", Value: 10733223936, Help: "PSize - Size of PV in current units."},
						{Field: "pv_used", Value: 3225419776, Help: "Used - Total amount of allocated space in current units."},
					},
					Labels: []Label{
						{Field: "pv_name", Value: "/dev/sdc"},
						{Field: "pv_uuid", Value: "PTjDUr-PO0k-kqFA-99Nc-qqsM-bBe0-tHxm5q"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReportsToMetrics(tt.args.b)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
