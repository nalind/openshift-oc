package create

import (
	"context"
	"fmt"
	"strings"

	ocmdhelpers "github.com/openshift/oc/pkg/helpers/cmd"
	"github.com/spf13/cobra"
	"k8s.io/client-go/discovery"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util"
	"k8s.io/kubectl/pkg/util/templates"

	userv1 "github.com/openshift/api/user/v1"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

var (
	userLong = templates.LongDesc(`
		This command can be used to create a user object.

		Typically, users are created automatically during login. If automatic
		creation is disabled (by using the "lookup" mapping method), users must
		be created manually.

		Corresponding identity and user identity mapping objects must also be created
		to allow logging in as the created user.
	`)

	userExample = templates.Examples(`
		# Create a user with the username "ajones" and the display name "Adam Jones"
		oc create user ajones --full-name="Adam Jones"
	`)
)

type CreateUserOptions struct {
	CreateSubcommandOptions *CreateSubcommandOptions

	FullName string

	UserClient      userv1client.UsersGetter
	DiscoveryClient discovery.DiscoveryInterface
}

// NewCmdCreateUser is a macro command to create a new user
func NewCmdCreateUser(f genericclioptions.RESTClientGetter, streams genericiooptions.IOStreams) *cobra.Command {
	o := &CreateUserOptions{
		CreateSubcommandOptions: NewCreateSubcommandOptions(streams),
	}
	cmd := &cobra.Command{
		Use:     "user NAME",
		Short:   "Manually create a user (only needed if automatic creation is disabled)",
		Long:    userLong,
		Example: userExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd, f, args))
			ocmdhelpers.CheckOAuthDisabledErr(o.Run(), o.DiscoveryClient)
		},
	}
	cmd.Flags().StringVar(&o.FullName, "full-name", o.FullName, "Display name of the user")

	o.CreateSubcommandOptions.AddFlags(cmd)
	cmdutil.AddDryRunFlag(cmd)

	return cmd
}

func (o *CreateUserOptions) Complete(cmd *cobra.Command, f genericclioptions.RESTClientGetter, args []string) error {
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	o.UserClient, err = userv1client.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	o.DiscoveryClient, err = f.ToDiscoveryClient()
	if err != nil {
		return err
	}

	return o.CreateSubcommandOptions.Complete(f, cmd, args)
}

func (o *CreateUserOptions) Run() error {
	if strings.ContainsAny(o.CreateSubcommandOptions.Name, "\r\n") {
		return fmt.Errorf("new line in user name is not allowed")
	}

	user := &userv1.User{
		// this is ok because we know exactly how we want to be serialized
		TypeMeta: metav1.TypeMeta{APIVersion: userv1.SchemeGroupVersion.String(), Kind: "User"},
		ObjectMeta: metav1.ObjectMeta{
			Name: o.CreateSubcommandOptions.Name,
		},
		FullName: o.FullName,
	}

	if err := util.CreateOrUpdateAnnotation(o.CreateSubcommandOptions.CreateAnnotation, user, createCmdJSONEncoder()); err != nil {
		return err
	}

	var err error
	if o.CreateSubcommandOptions.DryRunStrategy != cmdutil.DryRunClient {
		user, err = o.UserClient.Users().Create(context.TODO(), user, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return o.CreateSubcommandOptions.Printer.PrintObj(user, o.CreateSubcommandOptions.Out)
}
