package awr

const AwrTpl = `
<html lang="en">
<head>
	<title>AWR Report for PostgreSQL</title>
	<style type="text/css">
		body.awr {font:bold 10pt Arial,Helvetica,Geneva,sans-serif;color:black; background:White;}
		pre.awr  {font:8pt Courier;color:black; background:White;}
		h1.awr   {font:bold 20pt Arial,Helvetica,Geneva,sans-serif;color:#336699;background-color:White;border-bottom:1px solid #cccc99;margin-top:0pt; margin-bottom:0pt;padding:0px 0px 0px 0px;}
		h2.awr   {font:bold 18pt Arial,Helvetica,Geneva,sans-serif;color:#336699;background-color:White;margin-top:4pt; margin-bottom:0pt;}
		h3.awr {font:bold 16pt Arial,Helvetica,Geneva,sans-serif;color:#336699;background-color:White;margin-top:4pt; margin-bottom:0pt;}
		li.awr {font: 8pt Arial,Helvetica,Geneva,sans-serif; color:black; background:White;}
		th.awrnobg {font:bold 8pt Arial,Helvetica,Geneva,sans-serif; color:black; background:White;padding-left:4px; padding-right:4px;padding-bottom:2px}
		th.awrbg {font:bold 8pt Arial,Helvetica,Geneva,sans-serif; color:White; background:#0066CC;padding-left:4px; padding-right:4px;padding-bottom:2px}
		td.awrnc {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;}
		td.awrc    {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;}
		td.awrnclb {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-left: thin solid black;}
		td.awrncbb {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-left: thin solid black;border-right: thin solid black;}
		td.awrncrb {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-right: thin solid black;}
		td.awrcrb    {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-right: thin solid black;}
		td.awrclb    {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-left: thin solid black;}
		td.awrcbb    {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-left: thin solid black;border-right: thin solid black;}
		a.awr {font:bold 8pt Arial,Helvetica,sans-serif;color:#663300; vertical-align:top;margin-top:0pt; margin-bottom:0pt;}
		td.awrnct {font:8pt Arial,Helvetica,Geneva,sans-serif;border-top: thin solid black;color:black;background:White;vertical-align:top;}
		td.awrct   {font:8pt Arial,Helvetica,Geneva,sans-serif;border-top: thin solid black;color:black;background:#FFFFCC; vertical-align:top;}
		td.awrnclbt  {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-top: thin solid black;border-left: thin solid black;}
		td.awrncbbt  {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-left: thin solid black;border-right: thin solid black;border-top: thin solid black;}
		td.awrncrbt {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:White;vertical-align:top;border-top: thin solid black;border-right: thin solid black;}
		td.awrcrbt     {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-top: thin solid black;border-right: thin solid black;}
		td.awrclbt     {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-top: thin solid black;border-left: thin solid black;}
		td.awrcbbt   {font:8pt Arial,Helvetica,Geneva,sans-serif;color:black;background:#FFFFCC; vertical-align:top;border-top: thin solid black;border-left: thin solid black;border-right: thin solid black;}
		table.tdiff {  border_collapse: collapse; }
		.hidden   {position:absolute;left:-10000px;top:auto;width:1px;height:1px;overflow:hidden;}
		.pad   {margin-left:17px;}
		.doublepad {margin-left:34px;}
	</style>
</head>
<body class="awr">
	<h1 class="awr">PostgreSQL AWR</h1>
	<h2 class="awr"><a class="awr" name="99999"></a>System Information</h2>
	<p>
	<table border="0" width="800" class="tdiff" summary="This table displays system information">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">Hostname</th>
                <th class="awrbg" scope="col">RunTime</th>
                <th class="awrbg" scope="col">CPU</th>
                <th class="awrbg" scope="col">Load Average</th>
                <th class="awrbg" scope="col">Memory</th>
                <th class="awrbg" scope="col">Swap</th>
			</tr>
            <tr>
                <td scope="row" class="awrnc">{{ .SystemInfo.Hostname }}</td>
				<td scope="row" class="awrnc">{{ .SystemInfo.RunTime }}</td>
				<td scope="row" class="awrnc">{{ .SystemInfo.Cpu }}</td>
				<td scope="row" class="awrnc">{{ .SystemInfo.LoadAvg }}</td>
				<td scope="row" class="awrnc">{{ .SystemInfo.Memory }}</td>
				<td scope="row" class="awrnc">{{ .SystemInfo.Swap }}</td>
            </tr>
        </tbody>
    </table>
	</p>
	<h2 class="awr"><a class="awr" name="99999"></a>PostgreSQL Cluster Information</h2>
	<p>
	<table border="0" width="800" class="tdiff" summary="This table displays postgresql cluster information">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">Version</th>
                <th class="awrbg" scope="col">StartTime</th>
                <th class="awrbg" scope="col">SystemIdentifier</th>
                <th class="awrbg" scope="col">TimeLine</th>
                <th class="awrbg" scope="col">LastCheckpointTime</th>
                <th class="awrbg" scope="col">RedoLsn</th>
				<th class="awrbg" scope="col">CheckpointLsn</th>
				<th class="awrbg" scope="col">TransactionId</th>
				<th class="awrbg" scope="col">RecoveryMode?</th>
			</tr>
            <tr>
                <td scope="row" class="awrnc">{{ .PgClusterInfo.ServerVersion }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.ServerStartTime }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.SystemIdentifier }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.TimeLine }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.LastCheckpointTime }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.RedoLsn }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.CheckpointLsn }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.TransactionId }}</td>
				<td scope="row" class="awrnc">{{ .PgClusterInfo.RecoveryMode }}</td>
            </tr>
        </tbody>
    </table>
	</p>
	<h2 class="awr"><a class="awr" name="99999"></a>PostgreSQL Cluster Prameter</h2>
	<p>
	<table border="0" width="800" class="tdiff" summary="This table displays postgresql cluster parameter">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">name</th>
                <th class="awrbg" scope="col">value</th>
			</tr>
			{{ range $index, $item := .PgClusterInfo.ParamSet.ParamItems }}
			{{ if eq (mod $index 2) 1 }}
            <tr>
				<td scope="row" class="awrc">{{ $item.Name }}</td>
				<td scope="row" class="awrc">{{ $item.Value }}</td>
            </tr>
			{{ end }}
			{{ if eq (mod $index 2) 0 }}
            <tr>
				<td scope="row" class="awrnc">{{ $item.Name }}</td>
				<td scope="row" class="awrnc">{{ $item.Value }}</td>
            </tr>
			{{ end }}
			{{ end }}
        </tbody>
    </table>
	</p>
	<h2 class="awr">SQL Statistics</h2>
	<ul>
		<li class="awr"><a class="awr" href="#550">SQL ordered by User I/O</a></li>
		<li class="awr"><a class="awr" href="#560">SQL ordered by Elapsed Time</a></li>
		<li class="awr"><a class="awr" href="#570">SQL ordered by Shared Buffer</a></li>
		<li class="awr"><a class="awr" href="#580">SQL ordered by Temp</a></li>
	</ul>
	<p>
		<a class="awr" name="550"></a>
	</p>
	<h3 class="awr">SQL ordered by User I/O</h3>
	<table border="0" width="800" class="tdiff" summary="This table displays top SQL by User I/O">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">userid</th>
				<th class="awrbg" scope="col">dbid</th>
				<th class="awrbg" scope="col">calls</th>
				<th class="awrbg" scope="col">min_exec_time</th>
				<th class="awrbg" scope="col">max_exec_time</th>
				<th class="awrbg" scope="col">mean_exec_time</th>
				<th class="awrbg" scope="col">total_exec_time</th>
				<th class="awrbg" scope="col">shared_blks_hit</th>
				<th class="awrbg" scope="col">shared_blks_read</th>
				<th class="awrbg" scope="col">shared_blks_dirtied</th>
				<th class="awrbg" scope="col">shared_blks_written</th>
				<th class="awrbg" scope="col">temp_blks_read</th>
				<th class="awrbg" scope="col">temp_blks_written</th>
				<th class="awrbg" scope="col">blk_read_time</th>
				<th class="awrbg" scope="col">blk_write_time</th>
				<th class="awrbg" scope="col">query</th>
			</tr>
			{{ range $index, $item := .ComsumeIoSqlSet.ComsumeTopSqlItems }}
			{{ if eq (mod $index 2) 1 }}
            <tr>
				<td scope="row" class="awrc">{{ $item.UserId }}</td>
				<td scope="row" class="awrc">{{ $item.DbId }}</td>
				<td scope="row" class="awrc">{{ $item.Calls }}</td>
				<td scope="row" class="awrc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ if eq (mod $index 2) 0 }}
            <tr>
				<td scope="row" class="awrnc">{{ $item.UserId }}</td>
				<td scope="row" class="awrnc">{{ $item.DbId }}</td>
				<td scope="row" class="awrnc">{{ $item.Calls }}</td>
				<td scope="row" class="awrnc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrnc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ end }}
        </tbody>
    </table>
	<p>
		<a class="awr" name="560"></a>
	</p>
	<h3 class="awr">SQL ordered by Elapsed Time</h3>
	<table border="0" width="800" class="tdiff" summary="This table displays top SQL by Elapsed Time">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">userid</th>
				<th class="awrbg" scope="col">dbid</th>
				<th class="awrbg" scope="col">calls</th>
				<th class="awrbg" scope="col">min_exec_time</th>
				<th class="awrbg" scope="col">max_exec_time</th>
				<th class="awrbg" scope="col">mean_exec_time</th>
				<th class="awrbg" scope="col">total_exec_time</th>
				<th class="awrbg" scope="col">shared_blks_hit</th>
				<th class="awrbg" scope="col">shared_blks_read</th>
				<th class="awrbg" scope="col">shared_blks_dirtied</th>
				<th class="awrbg" scope="col">shared_blks_written</th>
				<th class="awrbg" scope="col">temp_blks_read</th>
				<th class="awrbg" scope="col">temp_blks_written</th>
				<th class="awrbg" scope="col">blk_read_time</th>
				<th class="awrbg" scope="col">blk_write_time</th>
				<th class="awrbg" scope="col">query</th>
			</tr>
			{{ range $index, $item := .ComsumeTimeSqlSet.ComsumeTopSqlItems }}
			{{ if eq (mod $index 2) 1 }}
            <tr>
				<td scope="row" class="awrc">{{ $item.UserId }}</td>
				<td scope="row" class="awrc">{{ $item.DbId }}</td>
				<td scope="row" class="awrc">{{ $item.Calls }}</td>
				<td scope="row" class="awrc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ if eq (mod $index 2) 0 }}
            <tr>
				<td scope="row" class="awrnc">{{ $item.UserId }}</td>
				<td scope="row" class="awrnc">{{ $item.DbId }}</td>
				<td scope="row" class="awrnc">{{ $item.Calls }}</td>
				<td scope="row" class="awrnc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrnc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ end }}
        </tbody>
    </table>
	<p>
		<a class="awr" name="570"></a>
	</p>
	<h3 class="awr">SQL ordered by Shared Buffer</h3>
	<table border="0" width="800" class="tdiff" summary="This table displays top SQL by Shared Buffer">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">userid</th>
				<th class="awrbg" scope="col">dbid</th>
				<th class="awrbg" scope="col">calls</th>
				<th class="awrbg" scope="col">min_exec_time</th>
				<th class="awrbg" scope="col">max_exec_time</th>
				<th class="awrbg" scope="col">mean_exec_time</th>
				<th class="awrbg" scope="col">total_exec_time</th>
				<th class="awrbg" scope="col">shared_blks_hit</th>
				<th class="awrbg" scope="col">shared_blks_read</th>
				<th class="awrbg" scope="col">shared_blks_dirtied</th>
				<th class="awrbg" scope="col">shared_blks_written</th>
				<th class="awrbg" scope="col">temp_blks_read</th>
				<th class="awrbg" scope="col">temp_blks_written</th>
				<th class="awrbg" scope="col">blk_read_time</th>
				<th class="awrbg" scope="col">blk_write_time</th>
				<th class="awrbg" scope="col">query</th>
			</tr>
			{{ range $index, $item := .ComsumeBufferSqlSet.ComsumeTopSqlItems }}
			{{ if eq (mod $index 2) 1 }}
            <tr>
				<td scope="row" class="awrc">{{ $item.UserId }}</td>
				<td scope="row" class="awrc">{{ $item.DbId }}</td>
				<td scope="row" class="awrc">{{ $item.Calls }}</td>
				<td scope="row" class="awrc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ if eq (mod $index 2) 0 }}
            <tr>
				<td scope="row" class="awrnc">{{ $item.UserId }}</td>
				<td scope="row" class="awrnc">{{ $item.DbId }}</td>
				<td scope="row" class="awrnc">{{ $item.Calls }}</td>
				<td scope="row" class="awrnc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrnc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ end }}
        </tbody>
    </table>
	<p>
		<a class="awr" name="580"></a>
	</p>
	<h3 class="awr">SQL ordered by Temp</h3>
	<table border="0" width="800" class="tdiff" summary="This table displays top SQL by Temp">
        <tbody>
            <tr>
                <th class="awrbg" scope="col">userid</th>
				<th class="awrbg" scope="col">dbid</th>
				<th class="awrbg" scope="col">calls</th>
				<th class="awrbg" scope="col">min_exec_time</th>
				<th class="awrbg" scope="col">max_exec_time</th>
				<th class="awrbg" scope="col">mean_exec_time</th>
				<th class="awrbg" scope="col">total_exec_time</th>
				<th class="awrbg" scope="col">shared_blks_hit</th>
				<th class="awrbg" scope="col">shared_blks_read</th>
				<th class="awrbg" scope="col">shared_blks_dirtied</th>
				<th class="awrbg" scope="col">shared_blks_written</th>
				<th class="awrbg" scope="col">temp_blks_read</th>
				<th class="awrbg" scope="col">temp_blks_written</th>
				<th class="awrbg" scope="col">blk_read_time</th>
				<th class="awrbg" scope="col">blk_write_time</th>
				<th class="awrbg" scope="col">query</th>
			</tr>
			{{ range $index, $item := .ComsumeTempSqlSet.ComsumeTopSqlItems }}
			{{ if eq (mod $index 2) 1 }}
            <tr>
				<td scope="row" class="awrc">{{ $item.UserId }}</td>
				<td scope="row" class="awrc">{{ $item.DbId }}</td>
				<td scope="row" class="awrc">{{ $item.Calls }}</td>
				<td scope="row" class="awrc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ if eq (mod $index 2) 0 }}
            <tr>
				<td scope="row" class="awrnc">{{ $item.UserId }}</td>
				<td scope="row" class="awrnc">{{ $item.DbId }}</td>
				<td scope="row" class="awrnc">{{ $item.Calls }}</td>
				<td scope="row" class="awrnc">{{ $item.MinExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MaxExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.MeanExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.TotalExecTime }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksHit }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksDirtied }}</td>
				<td scope="row" class="awrnc">{{ $item.SharedBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksRead }}</td>
				<td scope="row" class="awrnc">{{ $item.TempBlksWritten }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkReadTime }}</td>
				<td scope="row" class="awrnc">{{ $item.BlkWriteTime }}</td>
				<td scope="row" class="awrnc">{{ $item.Query }}</td>
            </tr>
			{{ end }}
			{{ end }}
        </tbody>
    </table>
</body>
</html>
`