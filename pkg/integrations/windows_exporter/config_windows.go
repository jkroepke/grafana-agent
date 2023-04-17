package windows_exporter //nolint:golint

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus-community/windows_exporter/collector"
	"time"
)

const DefaultCollectors = "cpu,cs,logical_disk,net,os,service,system,textfile"

// Populate defaults for all collector configs.
func init() {
	// Register flags from all collector configs to a fake integration and then
	// parse an empty command line to force defaults to be populated.
	app := kingpin.New("", "")
	_, err := app.Parse([]string{})
	if err != nil {
		panic(err)
	}

	// Map the configs with defaults applied to our default config.
	DefaultConfig.fromExporterConfig(app)
}

// fromExporterConfig converts windows_exporter configs into the integration Config.
func (c *Config) fromExporterConfig(app *kingpin.Application) {
	c.EnabledCollectors = DefaultCollectors
	c.Timeout = 4 * time.Minute

	c.Dfsr.SourcesEnabled = *app.GetFlag(collector.FlagDfsrEnabledCollectors).String()
	c.Exchange.EnabledList = *app.GetFlag(collector.FlagExchangeCollectorsEnabled).String()
	c.IIS.SiteExclude = *app.GetFlag(collector.FlagIISAppExclude).String()
	c.IIS.SiteInclude = *app.GetFlag(collector.FlagIISSiteInclude).String()
	c.IIS.AppExclude = *app.GetFlag(collector.FlagIISAppExclude).String()
	c.IIS.AppInclude = *app.GetFlag(collector.FlagIISAppInclude).String()
	c.LogicalDisk.Exclude = *app.GetFlag(collector.FlagLogicalDiskVolumeExclude).String()
	c.LogicalDisk.Include = *app.GetFlag(collector.FlagLogicalDiskVolumeInclude).String()
	c.MSMQ.Where = *app.GetFlag(collector.FlagMsmqWhereClause).String()
	c.MSSQL.EnabledClasses = *app.GetFlag(collector.FlagMssqlEnabledCollectors).String()
	c.Network.Exclude = *app.GetFlag(collector.FlagNicExclude).String()
	c.Network.Include = *app.GetFlag(collector.FlagNicInclude).String()
	c.Process.Exclude = *app.GetFlag(collector.FlagProcessExclude).String()
	c.Process.Include = *app.GetFlag(collector.FlagProcessInclude).String()
	c.ScheduledTask.Exclude = *app.GetFlag(collector.FlagScheduledTaskExclude).String()
	c.ScheduledTask.Include = *app.GetFlag(collector.FlagScheduledTaskInclude).String()
	c.Service.Where = *app.GetFlag(collector.FlagServiceWhereClause).String()
	c.Service.UseApi = *app.GetFlag(collector.FlagServiceUseAPI).String()
	c.SMTP.Exclude = *app.GetFlag(collector.FlagSmtpServerExclude).String()
	c.SMTP.Include = *app.GetFlag(collector.FlagSmtpServerInclude).String()
	c.TextFile.TextFileDirectory = *app.GetFlag(collector.FlagTextFileDirectory).String()
}

// toExporterConfig converts integration Configs into windows_exporter configs.
func (c *Config) toExporterConfig(app *kingpin.Application) {
	app.GetFlag(collector.FlagDfsrEnabledCollectors).StringVar(&c.Dfsr.SourcesEnabled)
	app.GetFlag(collector.FlagExchangeCollectorsEnabled).StringVar(&c.Exchange.EnabledList)
	app.GetFlag(collector.FlagIISSiteExclude).StringVar(&c.IIS.SiteExclude)
	app.GetFlag(collector.FlagIISSiteInclude).StringVar(&c.IIS.SiteInclude)
	app.GetFlag(collector.FlagIISAppExclude).StringVar(&c.IIS.AppExclude)
	app.GetFlag(collector.FlagIISAppInclude).StringVar(&c.IIS.AppInclude)
	app.GetFlag(collector.FlagLogicalDiskVolumeExclude).StringVar(&c.LogicalDisk.Exclude)
	app.GetFlag(collector.FlagLogicalDiskVolumeInclude).StringVar(&c.LogicalDisk.Include)
	app.GetFlag(collector.FlagMsmqWhereClause).StringVar(&c.MSMQ.Where)
	app.GetFlag(collector.FlagMssqlEnabledCollectors).StringVar(&c.MSSQL.EnabledClasses)
	app.GetFlag(collector.FlagNicExclude).StringVar(&c.Network.Exclude)
	app.GetFlag(collector.FlagNicInclude).StringVar(&c.Network.Include)
	app.GetFlag(collector.FlagProcessExclude).StringVar(&c.Process.Exclude)
	app.GetFlag(collector.FlagProcessInclude).StringVar(&c.Process.Include)
	app.GetFlag(collector.FlagScheduledTaskExclude).StringVar(&c.Process.Exclude)
	app.GetFlag(collector.FlagScheduledTaskInclude).StringVar(&c.Process.Include)
	app.GetFlag(collector.FlagServiceWhereClause).StringVar(&c.Service.Where)
	app.GetFlag(collector.FlagServiceUseAPI).StringVar(&c.Service.UseApi)
	app.GetFlag(collector.FlagSmtpServerExclude).StringVar(&c.SMTP.Exclude)
	app.GetFlag(collector.FlagNicInclude).StringVar(&c.SMTP.Include)
	app.GetFlag(collector.FlagTextFileDirectory).StringVar(&c.TextFile.TextFileDirectory)
}
