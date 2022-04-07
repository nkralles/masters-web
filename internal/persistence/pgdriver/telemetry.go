package pgdriver

import (
	"context"
	"github.com/nkralles/masters-web/internal/persistence"
)

func (d *Driver) HttpTelemetry(ctx context.Context, t persistence.Telemetry) {

	_, _ = d.pool.Exec(ctx, `insert into rest_telemetry(ip, http_method, url_path, http_code, http_written, http_duration, http_duration_text,
                           user_agent_family, user_agent_major, user_agent_minor, user_agent_patch, os_family, os_major,
                           os_minor, os_patch, os_patch_minor, device_family, device_brand, device_model)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);`,
		t.IP,
		t.HttpMethod,
		t.UrlPath,
		t.HttpCode,
		t.HttpWritten,
		t.HttpDuration.Seconds(),
		t.HttpDuration.String(),
		t.Ua.UserAgent.Family,
		t.Ua.UserAgent.Major,
		t.Ua.UserAgent.Minor,
		t.Ua.UserAgent.Patch,
		t.Ua.Os.Family,
		t.Ua.Os.Major,
		t.Ua.Os.Minor,
		t.Ua.Os.Patch,
		t.Ua.Os.PatchMinor,
		t.Ua.Device.Family,
		t.Ua.Device.Brand,
		t.Ua.Device.Model,
	)
}
