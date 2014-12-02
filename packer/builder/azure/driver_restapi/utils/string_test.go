// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package utils

import (
	"testing"
)

func TestClue(t *testing.T) {

	a := "Inline script!\n\nHandles  NPM(K)    PM(K)      WS(K) VM(M)   CPU(s)     Id ProcessName          \n-------  ------    -----      ----- -----   ------     -- -----------          \n     48       3     1436       1936    12     0.00   1956 cmd                  \n     39       3     1424       1928    12     0.02   2456 cmd                  \n     42       4      552       2520    23     0.02    952 conhost              \n     42       4      552       2492    23     0.02   2424 conhost              \n    217      12     1600       3468    43     0.20    364 csrss                \n    111       9     1292       3488    39     0.13    416 csrss                \n    508      38    22280      29376   538     0.77   2468 CustomScriptHandler  \n    176      14    14716      24552    87     0.11    688 dwm                  \n      0       0        0          4     0               0 Idle                 \n    285      21    11924      23636   133     0.13   2644 LogonUI              \n    712      18     3508       9184    36     0.72    516 lsass                \n    422      31    75416      77168   608     1.81    840 powershell           \n    121       9     2460       5144    58     0.02   2852 rundll32             \n    121       9     2676       5384    58     0.05   3016 rundll32             \n    226      11     2856       5928    21     3.75    508 services             \n     52       3      304       1052     5     0.08    260 smss                 \n    385      21     3304       8584    70     0.13   1216 spoolsv              \n    190       9     3988       9932    44     1.50    664 sppsvc               \n   1313      66    30044      43192   454     8.66    576 svchost              \n    372      15     3508       7808    38     0.22    580 svchost              \n    326      15     2268       5480    25     0.19    608 svchost              \n    565      34     6784      14680  1130     0.53    860 svchost              \n    407      21     4324       9436    69     0.25    900 svchost              \n    490      22    13784      16604    68     0.64    972 svchost              \n    356      31     5288      10424    48     0.25   1084 svchost              \n    262      14     3632       8552    46     2.53   1248 svchost              \n    179      17    11088       9884   323     0.03   1412 svchost              \n    131       9     1500       6428    35     0.06   1504 svchost              \n    266      16     2580       6932    46     0.06   1528 svchost              \n    416      19     3412       7656    73     0.31   1932 svchost              \n    612       0      116        280     3     3.41      4 System               \n    127       9     1332       5468    35     0.02   1688 VSSVC                \n    362      39    33124      41968   569     1.08   2512 WaAppAgent           \n    545      46    41564      54868   599     3.22   2216 WindowsAzureGuestA...\n    398      45    42376      51472   608     3.92   1028 WindowsAzureTeleme...\n     92       8      900       3560    21     0.08    424 wininit              \n    132       8     1312       8544    68     0.11    452 winlogon             \n    140      11     5776       9412    47     1.89   1856 WmiPrvSE             \n    277      17     4772      11192    71     0.64   1988 WmiPrvSE             \n    146       9     1880       5168    32     0.02   2476 WmiPrvSE             \nInstalling notepad...\nInline script finished!\nStarting file script...\nExecuting [DateTime]::Now...\n\nDate        : 9/7/2014 12:00:00 AM\nDay         : 7\nDayOfWeek   : Sunday\nDayOfYear   : 250\nHour        : 1\nKind        : Local\nMillisecond : 489\nMinute      : 0\nMonth       : 9\nSecond      : 35\nTicks       : 635456484354890234\nTimeOfDay   : 01:00:35.4890234\nYear        : 2014\nDateTime    : Sunday, September 7, 2014 1:00:35 AM\n\nExecuting Install-WindowsFeature -Name \"XPS-Viewer\" -IncludeAllSubFeature"
	b := ": 9\nSecond      : 35\nTicks       : 635456484354890234\nTimeOfDay   : 01:00:35.4890234\nYear        : 2014\nDateTime    : Sunday, September 7, 2014 1:00:35 AM\n\nExecuting Install-WindowsFeature -Name \"XPS-Viewer\" -IncludeAllSubFeature\n\nSuccess       : True\nRestartNeeded : No\nFeatureResult : {XPS Viewer}\nExitCode      : Success\n\nInstalling ReportViewer...\nInstalling Git Bash...\nInstalling Mozilla Firefox...\nHello world 1!\nHello world 2!\nHello world 3!\nHello world 4!\nHello world 5!\n     48       3     1436       1936    12     0.00   1956 cmd                  \n     39       3     1424       1928    12     0.02   2456 cmd                  \n     42       4      552       2524    23     0.11    952 conhost              \n     42       4      552       2492    23     0.02   2424 conhost              \n    216      13     1604       4764    46     0.31    364 csrss                \n    108       9     1240       3472    38     0.13    416 csrss                \n    480      35    22504      30184   532     0.80   2468 CustomScriptHandler  \n    173      14    14656      24516    86     0.11    688 dwm                  \n      0       0        0          4     0               0 Idle                 \n    276      21    11844      23272   132     0.13   2644 LogonUI              \n    680      18     3476       9428    35     0.81    516 lsass                \n    164      12     2160       6372    41     0.05   2820 msdtc                \n    343      15     7704      13980    89     1.20   2760 msiexec              \n    726      45    99896     110232   637     6.56    840 powershell           \n    121       9     2392       1480    58     0.02   2852 rundll32             \n    119       9     2608       1440    58     0.05   3016 rundll32             \n    236      10     2756       5924    20     3.77    508 services             \n     52       2      272       1032     4     0.08    260 smss                 \n    365      20     3032       8520    69     0.13   1216 spoolsv              \n    168       8     3768       6400    42     1.50    664 sppsvc               \n   1189      57    27080      38812   420     9.47    576 svchost              \n    355      15     3444       7756    38     0.22    580 svchost              \n    346      15     2432       5640    25     0.23    608 svchost              \n    656      34     6712      15236  1129     0.56    860 svchost              \n    382      22     4436       9632    68     0.25    900 svchost              \n    471      21    15504      18376    65     1.11    972 svchost              \n    381      33     5960      11024    52     0.28   1084 svchost              \n     91       7     1120       4336    27     0.02   1144 svchost              \n    384      24     9212      12864   626     2.61   1248 svchost              \n    176      17    11140       9876   324     0.03   1412 svchost              \n    127       9     1500       6412    35     0.06   1504 svchost              \n    257      16     2528       6924    46     0.08   1528 svchost              \n    398      18     3208       7552    71     0.31   1932 svchost              \n    643       0      116        280     3     4.02      4 System               \n    269      12     9992      15480    79     7.06   2068 TiWorker             \n    106       7     1412       4436    28     0.22   3008 TrustedInstaller     \n    129       9     1344       5492    35     0.02   1688 VSSVC                \n    378      38    32936      42540   568     1.30   2512 WaAppAgent           \n    538      44    41080      54636   595     3.34   2216 WindowsAzureGuestA...\n    390      43    42216      51412   603     3.92   1028 WindowsAzureTeleme...\n     83       8      824       3536    21     0.08    424 wininit              \n    126       8     1312       8544    68     0.11    452 winlogon             \n    128       9     6452       9348    45     1.89   1856 WmiPrvSE             \n    313      16     6520      13504    52     1.02   1988 WmiPrvSE             \nExecuting Sysprep from File..."

	a = Clue(a, b)

	t.Logf("a: '%s'", a)

	t.Error("eom")
}
