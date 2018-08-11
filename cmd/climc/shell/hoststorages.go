package shell

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
)

func init() {
	type HostStorageListOptions struct {
		BaseListOptions
		Host    string `help:"ID or Name of Host"`
		Storage string `help:"ID or Name of Storage"`
	}
	R(&HostStorageListOptions{}, "host-storage-list", "List host storage pairs", func(s *mcclient.ClientSession, args *HostStorageListOptions) error {
		params := FetchPagingParams(args.BaseListOptions)
		var result *modules.ListResult
		var err error
		if len(args.Storage) > 0 {
			params.Add(jsonutils.NewString(args.Storage), "storage")
		}
		if len(args.Host) > 0 {
			result, err = modules.Hoststorages.ListDescendent(s, args.Host, params)
		} else if len(args.Storage) > 0 {
			result, err = modules.Hoststorages.ListDescendent2(s, args.Storage, params)
		} else {
			result, err = modules.Hoststorages.List(s, params)
		}
		if err != nil {
			return err
		}
		printList(result, modules.Hoststorages.GetColumns(s))
		return nil
	})

	type HostStorageDetailOptions struct {
		HOST    string `help:"ID or Name of Host"`
		STORAGE string `help:"ID or Name of Storage"`
	}
	R(&HostStorageDetailOptions{}, "host-storage-show", "Show host storage details", func(s *mcclient.ClientSession, args *HostStorageDetailOptions) error {
		result, err := modules.Hoststorages.Get(s, args.HOST, args.STORAGE, nil)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&HostStorageDetailOptions{}, "host-storage-detach", "Detach a storage from a host", func(s *mcclient.ClientSession, args *HostStorageDetailOptions) error {
		result, err := modules.Hoststorages.Detach(s, args.HOST, args.STORAGE)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	type HostStorageAttachOptions struct {
		HOST       string `help:"ID or Name of Host"`
		STORAGE    string `help:"ID or Name of Storage"`
		MountPoint string `help:"MountPoint of Storage"`
	}

	R(&HostStorageAttachOptions{}, "host-storage-attach", "Attach a storage to a host", func(s *mcclient.ClientSession, args *HostStorageAttachOptions) error {
		params := jsonutils.NewDict()
		if args.MountPoint != "" {
			params.Add(jsonutils.NewString(args.MountPoint), "mount_point")
		}
		result, err := modules.Hoststorages.Attach(s, args.HOST, args.STORAGE, params)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

}
