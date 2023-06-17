#!/usr/bin/env python

data = """
Albania*AL*false*28*false*true*AL35202111090000000001234567
Andorra*AD*true*24*false*true*AD1400080001001234567890
Azerbaijan*AZ*false*28*false*false*AZ77VTBA00000000001234567890
Bahrain*BH*false*22*false*false*BH02CITI00001077181611
Brazil*BR*false*29*false*true*BR1500000000000010932840814P2
Bulgaria*BG*true*22*false*true*BG18RZBB91550123456789
Cyprus*CY*true*28*false*true*CY21002001950000357001234567
Dominican Republic*DO*false*28*false*false*DO22ACAU00000000000123456789
Egypt*EG*false*29*false*false*EG800002000156789012345180002
Georgia*GE*false*22*false*false*GE60NB0000000123456789
Gibraltar*GI*true*23*false*true*GI56XAPO000001234567890
Greece*GR*true*27*false*true*GR9608100010000001234567890
Guatemala*GT*false*28*false*false*GT20AGRO00000000001234567890
Jordan*JO*false*30*false*true*JO71CBJO0000000000001234567890
Kazakhstan*KZ*false*20*false*false*KZ244350000012344567
Kuwait*KW*false*30*false*false*KW81CBKU0000000000001234560101
Latvia*LV*true*21*false*false*LV97HABA0012345678910
Lebanon*LB*false*28*false*false*LB92000700000000123123456123
Lithuania*LT*true*20*false*true*LT601010012345678901
Luxembourg*LU*true*20*false*false*LU120010001234567891
Malta*MT*true*31*false*true*MT31MALT01100000000000000000123
Mauritius*MU*false*30*false*true*MU43BOMM0101123456789101000MUR
Moldova*MD*false*24*false*false*MD21EX000000000001234567
Pakistan*PK*false*24*false*false*PK36SCBL0000001123456702
Palestine*PS*false*29*false*false*PS92PALS000000000400123456702
Qatar*QA*false*29*false*false*QA54QNBA000000000000693123456
Romania*RO*true*24*false*false*RO66BACX0000001234567890
Saint Lucia*LC*false*32*false*false*LC14BOSL123456789012345678901234
Sao Tome and Principe*ST*false*25*false*false*ST23000200000289355710148
Saudi Arabia*SA*false*24*false*false*SA4420000001234567891234
Turkey*TR*false*26*false*true*TR320010009999901234567890
United Arab Emirates*AE*false*23*false*false*AE460090000000123456789
Holy See*VA*true*22*false*false*VA59001123000012345678
Virgin Islands, British*VG*false*24*false*false*VG07ABVI0000000123456789
Ukraine*UA*false*29*false*true*UA903052992990004149123456789
Seychelles*SC*false*31*false*false*SC74MCBL01031234567890123456USD
Iraq*IQ*false*23*false*false*IQ20CBIQ861800101010500
Belarus*BY*false*28*false*false*BY86AKBB10100000002966000000
El Salvador*SV*false*28*false*false*SV43ACAT00000000000000123123
Libya*LY*false*25*false*false*LY38021001000000123456789
Sudan*SD*false*18*false*false*SD8811123456789012
Burundi*BI*false*27*false*false*BI43220001131012345678912345
Djibouti*DJ*false*27*false*false*DJ2110002010010409943020008
Somalia*SO*false*23*false*true*SO061000001123123456789
Austria*AT*true*20*true*true*AT483200000012345864
Belgium*BE*true*16*true*true*BE71096123456769
Bosnia and Herzegovina*BA*false*20*true*true*BA393385804800211234
Costa Rica*CR*false*22*true*false*CR23015108410026012345
Croatia*HR*true*21*true*false*HR1723600001101234565
Czech Republic*CZ*true*24*true*false*CZ5508000000001234567899
Faroe Islands*FO*false*18*true*true*FO9264600123456789
Greenland*GL*false*18*true*true*GL8964710123456789
Denmark*DK*true*18*true*true*DK9520000123456789
Estonia*EE*true*20*true*true*EE471000001020145685
Finland*FI*true*18*true*true*FI1410093000123458
France*FR*true*27*true*true*FR7630006000011234567890189
Germany*DE*true*22*true*true*DE75512108001245126199
Hungary*HU*true*28*true*true*HU93116000060000000012345676
Iceland*IS*true*26*true*true*IS750001121234563108962099
Ireland*IE*true*22*true*true*IE64IRCE92050112345678
Israel*IL*false*23*true*true*IL170108000000012612345
Italy*IT*true*27*true*true*IT60X0542811101000000123456
Kosovo*XK*false*20*true*true*XK051212012345678906
Liechtenstein*LI*true*21*true*true*LI7408806123456789012
North Macedonia*MK*false*19*true*false*MK07200002785123453
Mauritania*MR*false*27*true*true*MR1300020001010000123456753
Monaco*MC*true*27*true*true*MC5810096180790123456789085
Montenegro*ME*false*22*true*false*ME25505000012345678951
Netherlands*NL*true*18*true*true*NL02ABNA0123456789
Norway*NO*true*15*true*true*NO8330001234567
Poland*PL*true*28*true*true*PL10105000997603123456789123
Portugal*PT*true*25*true*true*PT50002700000001234567833
San Marino*SM*true*27*true*true*SM76P0854009812123456789123
Serbia*RS*false*22*true*false*RS35105008123123123173
Slovakia*SK*true*24*true*false*SK8975000000000012345671
Slovenia*SI*true*19*true*true*SI56192001234567892
Spain*ES*true*24*true*true*ES7921000813610123456789
Sweden*SE*true*24*true*true*SE7280000810340009783242
Switzerland*CH*true*21*true*true*CH5604835012345678009
Timor-Leste*TL*false*23*true*false*TL380010012345678910106
Tunisia*TN*false*24*true*true*TN5904018104004942712345
United Kingdom*GB*true*22*true*true*GB33BUKB20201555555555
Russia*RU*false*33*true*true*RU0204452560040702810412345678901
"""


lines = data.splitlines()
for l in lines:
    chunks= l.split('*')
    print("{{\"{}\", \"{}\", true}},".format(chunks[0], chunks[-1]))
