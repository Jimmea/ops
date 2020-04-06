package lepton

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vmware/govmomi/vim25/methods"

	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
)

// Vsphere provides access to the Vsphere API.
type Vsphere struct {
	Storage *Objects
}

// BuildImage to be upload on v
func (v *Vsphere) BuildImage(ctx *Context) (string, error) {
	c := ctx.config
	err := BuildImage(*c)
	if err != nil {
		return "", err
	}

	return v.customizeImage(ctx)
}

// BuildImageWithPackage to upload on Vsphere.
func (v *Vsphere) BuildImageWithPackage(ctx *Context, pkgpath string) (string, error) {
	c := ctx.config
	err := BuildImageFromPackage(pkgpath, *c)
	if err != nil {
		return "", err
	}
	return v.customizeImage(ctx)
}

func (v *Vsphere) createImage(key string, bucket string, region string) {
}

// Initialize Vsphere related things
func (v *Vsphere) Initialize() error {
	return nil
}

// CreateImage - Creates image on vsphere using nanos images
func (v *Vsphere) CreateImage(ctx *Context) error {
	return nil
}

// ListImages lists images on Digital Ocean
func (v *Vsphere) ListImages(ctx *Context) error {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ID", "Status", "Created"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})
	table.SetRowLine(true)

	/*
		for _, image := range data {
			var row []string
			row = append(row, image.Description)
			row = append(row, image.SnapShotID)
			row = append(row, image.Status)
			row = append(row, image.CreatedAt)
			table.Append(row)
		}
	*/

	table.Render()

	return nil
}

// DeleteImage deletes image from v
func (v *Vsphere) DeleteImage(ctx *Context, imagename string) error {
	return nil
}

// CreateInstance - Creates instance on Digital Ocean Platform
func (v *Vsphere) CreateInstance(ctx *Context) error {
	return nil
}

// ListInstances lists instances on v
func (v *Vsphere) ListInstances(ctx *Context) error {

	venv := os.Getenv("GOVC_URL")

	u, err := url.Parse("https://" + venv + "/sdk")
	if err != nil {
		fmt.Println(err)
	}

	soapClient := soap.NewClient(u, true)
	c, err := vim25.NewClient(context.Background(), soapClient)
	if err != nil {
		fmt.Println(err)
	}

	req := types.Login{
		This: *c.ServiceContent.SessionManager,
	}

	req.UserName = u.User.Username()
	if pw, ok := u.User.Password(); ok {
		req.Password = pw
	}

	_, err = methods.Login(context.Background(), c, &req)
	if err != nil {
		fmt.Println(err)
	}

	// Create a view of HostSystem objects
	m := view.NewManager(c)

	v2, err := m.CreateContainerView(context.TODO(), c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer v2.Destroy(context.TODO())

	// Retrieve summary property for all hosts
	// Reference:
	// http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.HostSystem.html
	var hss []mo.HostSystem
	err = v2.Retrieve(context.TODO(), []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		return err
	}

	// Print summary per host (see also: govc/host/info.go)

	fmt.Println("woot?")
	fmt.Printf("%+v", hss)
	/*
		tw := tablewriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
		fmt.Fprintf(tw, "Name:\tUsed CPU:\tTotal CPU:\tFree CPU:\tUsed Memory:\tTotal Memory:\tFree Memory:\t\n")

		for _, hs := range hss {
			totalCPU := int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores)
			freeCPU := int64(totalCPU) - int64(hs.Summary.QuickStats.OverallCpuUsage)
			freeMemory := int64(hs.Summary.Hardware.MemorySize) - (int64(hs.Summary.QuickStats.OverallMemoryUsage) * 1024 * 1024)
			fmt.Fprintf(tw, "%s\t", hs.Summary.Config.Name)
			fmt.Fprintf(tw, "%d\t", hs.Summary.QuickStats.OverallCpuUsage)
			fmt.Fprintf(tw, "%d\t", totalCPU)
			fmt.Fprintf(tw, "%d\t", freeCPU)
			fmt.Fprintf(tw, "%s\t", hs.Summary.QuickStats.OverallMemoryUsage)
			fmt.Fprintf(tw, "%s\t", hs.Summary.Hardware.MemorySize)
			fmt.Fprintf(tw, "%d\t", freeMemory)
			fmt.Fprintf(tw, "\n")
		}

		_ = tw.Flush()
	*/
	return nil
}

// DeleteInstance deletes instance from v
func (v *Vsphere) DeleteInstance(ctx *Context, instancename string) error {
	return nil
}

// StartInstance starts an instance in v
func (v *Vsphere) StartInstance(ctx *Context, instancename string) error {
	return nil
}

// StopInstance deletes instance from v
func (v *Vsphere) StopInstance(ctx *Context, instancename string) error {
	return nil
}

// GetInstanceLogs gets instance related logs
func (v *Vsphere) GetInstanceLogs(ctx *Context, instancename string, watch bool) error {
	return nil
}

// TOv - make me shared
func (v *Vsphere) customizeImage(ctx *Context) (string, error) {
	imagePath := ctx.config.RunConfig.Imagename
	return imagePath, nil
}
