from bs4 import BeautifulSoup


# def parse_html(html_string):
#    soup = BeautifulSoup(html_string, "html.parser")
#    rows = soup.find_all('tr')
#    email_dict = {}
#    for row in rows:
#        if not row.th.string or not row.td.string:
#            continue
#        email = row.th.string.strip()
#        status = row.td.string.strip()
#        email_dict[email] = status == "Valid"
#    return email_dict

# https://www.rohannagar.com/jmail/
def parse_html(html_string):
    soup = BeautifulSoup(html_string, "html.parser")
    rows = soup.find_all('tr')
    email_dict = {}
    for row in rows:
        email = row.th.get_text().strip()  # Change here
        status = row.td.get_text().strip()  # And here
        email_dict[email] = status[0].lower() == "v"  # And here
    return email_dict


html_string = """
<tr>
      <th scope="row" valign="middle">"qu@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Opening quote must have a closing quote</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 12,521.84 ns</td>
      <td valign="middle">Invalid<br>in 84,668.56 ns</td>
      <td valign="middle">Invalid<br>in 37,805.72 ns</td>
      <td valign="middle">Invalid<br>in 65,059.67 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">ote"@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Closing quote must have an opening quote</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,375.52 ns</td>
      <td valign="middle">Invalid<br>in 44,958.71 ns</td>
      <td valign="middle">Invalid<br>in 41,474.26 ns</td>
      <td valign="middle">Invalid<br>in 41,550.87 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"(),:;&lt;&gt;[\]@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Opening quote must have a closing quote</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 15,959.25 ns</td>
      <td valign="middle">Invalid<br>in 23,460.15 ns</td>
      <td valign="middle">Invalid<br>in 59,579.83 ns</td>
      <td valign="middle">Invalid<br>in 60,757.67 ns</td>
    </tr>
    <tr>
      <th scope="row" valign="middle">\"\"\"@iana.org < br > </th >
      <td valign = "middle" > Invalid < br > <small class = "text-muted" > Each quote must be in a pair < /small > </td >
      <td style = "background-color:#56666B" valign = "middle" > Invalid < br > in 20, 420.3 ns < /td >
      <td valign = "middle" > Invalid < br > in 19, 022.68 ns < /td >
      <td valign = "middle" > Invalid < br > in 47, 801.29 ns < /td >
      <td valign = "middle" > Invalid < br > in 17, 438.2 ns < /td >
    </tr><tr>
      <th scope="row" valign="middle">a@b.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The high octet preset character is not allowed at the end of the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 89,513.7 ns</td>
      <td valign="middle">Invalid<br>in 350,849.96 ns</td>
      <td valign="middle">Invalid<br>in 37,629.76 ns</td>
      <td valign="middle">Invalid<br>in 14,591.56 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Abc.example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The @ character must be present</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,662.46 ns</td>
      <td valign="middle">Invalid<br>in 3,318.82 ns</td>
      <td valign="middle">Invalid<br>in 50,894.21 ns</td>
      <td valign="middle">Invalid<br>in 30,288.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">A@b@c@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There can only be a single @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,562.32 ns</td>
      <td valign="middle">Invalid<br>in 6,634.97 ns</td>
      <td valign="middle">Invalid<br>in 45,345.2 ns</td>
      <td valign="middle">Invalid<br>in 17,189.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a"b(c)d,e:f;g&lt;h&gt;i[j\k]l@example.co<br>m<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Certain characters cannot be in the local-part unquoted</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,569.64 ns</td>
      <td valign="middle">Invalid<br>in 4,878.58 ns</td>
      <td valign="middle">Invalid<br>in 65,591.46 ns</td>
      <td valign="middle">Invalid<br>in 12,004.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">just"not"right@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Quoted strings must be separated by dots</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,621.65 ns</td>
      <td valign="middle">Invalid<br>in 7,434.28 ns</td>
      <td valign="middle">Invalid<br>in 71,688.79 ns</td>
      <td valign="middle">Invalid<br>in 17,066.7 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">this is"not\allowed@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace should be quoted or dot-separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,208.28 ns</td>
      <td valign="middle">Invalid<br>in 7,932.7 ns</td>
      <td valign="middle">Invalid<br>in 56,258.84 ns</td>
      <td valign="middle">Invalid<br>in 27,182.08 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">this\ still\"not\\allowed@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Escaped whitespaces should still be quoted or dot-separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,106.89 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 34,395.97 ns</td>
      <td valign="middle">Invalid<br>in 38,188.09 ns</td>
      <td valign="middle">Invalid<br>in 17,464.57 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">1234567890123456789012345678901234567890<br>123456789012345678901234+x@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot have more than 64 characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 19,606.49 ns</td>
      <td valign="middle">Invalid<br>in 5,958.05 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 29,215.21 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 116,794.38 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">i_like_underscore@but_its_not_allowed_in<br>_this_part.example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Underscores are not allowed in the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 39,595.08 ns</td>
      <td valign="middle">Invalid<br>in 69,515.39 ns</td>
      <td valign="middle">Invalid<br>in 55,306.19 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 34,609.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">QA[icon]CHOCOLATE[icon]@test.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Unquoted [ and ] characters are not allowed</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,280.18 ns</td>
      <td valign="middle">Invalid<br>in 4,364.51 ns</td>
      <td valign="middle">Invalid<br>in 31,023.15 ns</td>
      <td valign="middle">Invalid<br>in 34,700.99 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">plainaddress<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The @ character must be present</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,256.17 ns</td>
      <td valign="middle">Invalid<br>in 1,327.29 ns</td>
      <td valign="middle">Invalid<br>in 31,262.8 ns</td>
      <td valign="middle">Invalid<br>in 19,513.53 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part must not be empty</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,051.61 ns</td>
      <td valign="middle">Invalid<br>in 1,145.38 ns</td>
      <td valign="middle">Invalid<br>in 35,471.64 ns</td>
      <td valign="middle">Invalid<br>in 12,608 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">@NotAnEmail<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part must not be empty</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,982.39 ns</td>
      <td valign="middle">Invalid<br>in 1,045.27 ns</td>
      <td valign="middle">Invalid<br>in 35,506.21 ns</td>
      <td valign="middle">Invalid<br>in 12,711.75 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">.email@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot begin with the . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 150.65 ns</td>
      <td valign="middle">Invalid<br>in 2,526.37 ns</td>
      <td valign="middle">Invalid<br>in 32,585.53 ns</td>
      <td valign="middle">Invalid<br>in 15,925.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email.@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot end with the . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,690.94 ns</td>
      <td valign="middle">Invalid<br>in 5,341.15 ns</td>
      <td valign="middle">Invalid<br>in 40,285.19 ns</td>
      <td valign="middle">Invalid<br>in 17,562.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email..email@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot contain two dots in a row</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,597.79 ns</td>
      <td valign="middle">Invalid<br>in 7,605.38 ns</td>
      <td valign="middle">Invalid<br>in 47,088.13 ns</td>
      <td valign="middle">Invalid<br>in 31,446.49 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@-example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot start with the - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,448.63 ns</td>
      <td valign="middle">Invalid<br>in 28,369.72 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,616.97 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 18,790.24 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@111.222.333.44444<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The top level domain cannot be all numeric</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 16,135.07 ns</td>
      <td valign="middle">Invalid<br>in 41,769.73 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,921.42 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 25,071.45 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example..com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot contain two dots in a row</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,024.25 ns</td>
      <td valign="middle">Invalid<br>in 22,368.43 ns</td>
      <td valign="middle">Invalid<br>in 66,535.92 ns</td>
      <td valign="middle">Invalid<br>in 58,254.22 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">this\ is"really"not\allowed@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The backslash escape must escape a special character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,337.19 ns</td>
      <td valign="middle">Invalid<br>in 14,124.9 ns</td>
      <td valign="middle">Invalid<br>in 66,942.63 ns</td>
      <td valign="middle">Invalid<br>in 21,589.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@sub.do,com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The , character is not allowed in the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 15,816.42 ns</td>
      <td valign="middle">Invalid<br>in 12,618.93 ns</td>
      <td valign="middle">Invalid<br>in 37,751.59 ns</td>
      <td valign="middle">Invalid<br>in 33,163.39 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@[12.34.44.56<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain literal must have the closing bracket</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,509.91 ns</td>
      <td valign="middle">Invalid<br>in 22,583.5 ns</td>
      <td valign="middle">Invalid<br>in 50,463.99 ns</td>
      <td valign="middle">Invalid<br>in 42,360.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@14.44.56.34]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain literal must have the opening bracket</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 18,825.15 ns</td>
      <td valign="middle">Invalid<br>in 27,723.64 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,501.57 ns</td>
      <td valign="middle">Invalid<br>in 45,571.25 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@[1.1.23.5f]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv4 format must be valid (disallowed character)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 13,697.95 ns</td>
      <td valign="middle">Invalid<br>in 24,126.56 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 12,889.67 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,414.49 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@[3.256.255.23]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv4 format must be valid (256 is above the IP range)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,609.23 ns</td>
      <td valign="middle">Invalid<br>in 26,205.57 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,761.16 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 26,605.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first"last"@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The final quote is invalid (missing pair and not dot-separated)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,969.3 ns</td>
      <td valign="middle">Invalid<br>in 12,284.77 ns</td>
      <td valign="middle">Invalid<br>in 71,975.1 ns</td>
      <td valign="middle">Invalid<br>in 29,047.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[.12.34.56.78]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv4 format must be valid (leading dot)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 52,897.87 ns</td>
      <td valign="middle">Invalid<br>in 12,593.8 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,986.2 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 28,764.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[12.34.56.789]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv4 format must be valid (789 is above the IP range)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,265.34 ns</td>
      <td valign="middle">Invalid<br>in 19,172.94 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,519.66 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 49,641.56 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[::12.34.56.78]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">IPv6 domain literal must have IPv6 tag</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,699.07 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 48,335.67 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 8,194.02 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 58,659.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">x@x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot be more than 255 characters in total</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 56,716.71 ns</td>
      <td valign="middle">Invalid<br>in 101,888.79 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 49,900.58 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 44,762.45 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"\"@iana.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A single backslash is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,661.79 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 12,834.84 ns</td>
      <td valign="middle">Invalid<br>in 39,457.56 ns</td>
      <td valign="middle">Invalid<br>in 17,461.97 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first\\@last@iana.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The extra @ must be quote-escaped</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,843.89 ns</td>
      <td valign="middle">Invalid<br>in 6,029.48 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,142.63 ns</td>
      <td valign="middle">Invalid<br>in 21,031.89 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot be missing</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,472.07 ns</td>
      <td valign="middle">Invalid<br>in 4,629.98 ns</td>
      <td valign="middle">Invalid<br>in 31,693.42 ns</td>
      <td valign="middle">Invalid<br>in 21,985.27 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@example.com
<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Invalid characters in the domain are not allowed</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 13,356.84 ns</td>
      <td valign="middle">Invalid<br>in 28,520.06 ns</td>
      <td valign="middle">Invalid<br>in 48,191.63 ns</td>
      <td valign="middle">Invalid<br>in 29,143.09 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:7777]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">IPv6 address does not have enough segments)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 9,615.92 ns</td>
      <td valign="middle">Invalid<br>in 66,939.62 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,991.2 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 30,436.58 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:7777:8888:9999]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">IPv6 address has too many segments)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 27,030.21 ns</td>
      <td valign="middle">Invalid<br>in 35,319.84 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 8,505.66 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 37,539.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222::3333::4444:5<br>555:6666]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (Only one set of :: is allowed)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 10,787 ns</td>
      <td valign="middle">Invalid<br>in 23,277.89 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 8,286.6 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 39,452.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:333x::4444:55<br>55]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (x is not a valid hexadecimal character)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,869.61 ns</td>
      <td valign="middle">Invalid<br>in 94,166.81 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 8,526.9 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 31,228.21 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:33333::4444:5<br>555]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (33333 is not a valid segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,947.46 ns</td>
      <td valign="middle">Invalid<br>in 60,886.39 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,199.7 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 25,227.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (: is invalid)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,198.25 ns</td>
      <td valign="middle">Invalid<br>in 58,612.84 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,669.53 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 18,722.88 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::::]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (::: is invalid)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,125.51 ns</td>
      <td valign="middle">Invalid<br>in 11,812.87 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,294.42 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 57,747.36 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,228.49 ns</td>
      <td valign="middle">Invalid<br>in 54,783.42 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,756.22 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 26,701.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::::b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,842.91 ns</td>
      <td valign="middle">Invalid<br>in 8,867.25 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,908.7 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 24,244.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::b3:b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,821.27 ns</td>
      <td valign="middle">Invalid<br>in 52,413.78 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,443.98 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 22,109.31 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::::b3:b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,985.36 ns</td>
      <td valign="middle">Invalid<br>in 14,818.55 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,069.73 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 24,102.56 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:::b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,276.37 ns</td>
      <td valign="middle">Invalid<br>in 23,034.77 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,627.08 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 30,998.4 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,673.69 ns</td>
      <td valign="middle">Invalid<br>in 9,077.93 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,955.28 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 21,937.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:::]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,980.37 ns</td>
      <td valign="middle">Invalid<br>in 8,199.74 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,377.24 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,400.25 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,489.64 ns</td>
      <td valign="middle">Invalid<br>in 13,489.95 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,197 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 26,398.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:::]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,527.16 ns</td>
      <td valign="middle">Invalid<br>in 15,464.06 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,793.06 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 41,811.52 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::b3:]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,350.13 ns</td>
      <td valign="middle">Invalid<br>in 9,205.69 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,923.98 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,680.72 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::a2::b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,709.63 ns</td>
      <td valign="middle">Invalid<br>in 10,071.03 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,165.46 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 28,582.73 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:a3:a4:b1:b2:b3:]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,545.81 ns</td>
      <td valign="middle">Invalid<br>in 16,351.06 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,044.78 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,335.89 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::a2:a3:a4:b1:b2:b3:b4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,779.02 ns</td>
      <td valign="middle">Invalid<br>in 16,746.95 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,058.34 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,543.11 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:a3:a4::b1:b2:b3:b<br>4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,031.6 ns</td>
      <td valign="middle">Invalid<br>in 13,341.77 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,826.87 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 29,473.7 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,253.49 ns</td>
      <td valign="middle">Invalid<br>in 57,642.03 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,535.44 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 29,707.09 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6::::11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,511.89 ns</td>
      <td valign="middle">Invalid<br>in 13,858.5 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,809.39 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,447.13 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,733.02 ns</td>
      <td valign="middle">Invalid<br>in 38,848.64 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 6,804.12 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 30,703.95 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:::11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,696.45 ns</td>
      <td valign="middle">Invalid<br>in 16,222.68 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,751.35 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 34,028.88 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:::11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (There is an extra colon after a2)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,953.97 ns</td>
      <td valign="middle">Invalid<br>in 16,565.4 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,175.1 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 31,879.42 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:cdef::11<br>.22.33.xx]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (IPv4 part has invalid characters)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 8,244.59 ns</td>
      <td valign="middle">Invalid<br>in 65,745.93 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 6,166.25 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 28,310.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv5:::12.34.56.78]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">IPv5 is not a valid domain literal</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,331.92 ns</td>
      <td valign="middle">Invalid<br>in 13,367.46 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,580.41 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 21,974.83 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:12.34.56.78]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (missing one IPv6 segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,087.89 ns</td>
      <td valign="middle">Invalid<br>in 38,234.94 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 6,464.58 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 25,258.25 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:7777:12.34.56.78]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (extra IPv6 segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,069.17 ns</td>
      <td valign="middle">Invalid<br>in 12,729.98 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,105.06 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 23,375.87 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:12.34.567.89]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (IPv4 part has an invalid segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,813.52 ns</td>
      <td valign="middle">Invalid<br>in 62,833.34 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,076.96 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 31,479.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">aaa@[123.123.123.123]a<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A domain literal cannot have extra characters outside of the brackets</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,981.77 ns</td>
      <td valign="middle">Invalid<br>in 22,112.04 ns</td>
      <td valign="middle">Invalid<br>in 38,094.82 ns</td>
      <td valign="middle">Invalid<br>in 70,555.54 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:CDEFF::1<br>1.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (CDEFF is not a valid segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,175.89 ns</td>
      <td valign="middle">Invalid<br>in 44,976.94 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,312.07 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 22,347.2 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::a4:b1::b4:11.22.33.<br>44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,243.79 ns</td>
      <td valign="middle">Invalid<br>in 16,233.75 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,229.64 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 24,955.15 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::11.22.33]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (IPv4 part is missing a segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,296.51 ns</td>
      <td valign="middle">Invalid<br>in 63,222.36 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,503.08 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 28,363.52 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::11.22.33.44.55]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (IPv4 part has an extra segment)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,964.17 ns</td>
      <td valign="middle">Invalid<br>in 62,608.88 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,194.47 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 27,047.84 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::b211.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (A colon is missing between the IPv6 and IPv4 parts)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,277.92 ns</td>
      <td valign="middle">Invalid<br>in 73,902.8 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,996.42 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 28,876.84 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::b2::11.22.33.44]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The IPv6 format is invalid (There is an extra colon before the IPv4 part)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,246.98 ns</td>
      <td valign="middle">Invalid<br>in 10,651.56 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,085.61 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 18,409.68 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a@-b.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot start with a - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 871.51 ns</td>
      <td valign="middle">Invalid<br>in 6,545.1 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,488.21 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,779.34 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a@b-.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A domain part cannot end with a - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 775.49 ns</td>
      <td valign="middle">Invalid<br>in 6,483.04 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 1,375.54 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,351.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">-@..com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be two . characters in a row in the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 662.1 ns</td>
      <td valign="middle">Invalid<br>in 5,389.76 ns</td>
      <td valign="middle">Invalid<br>in 29,968.49 ns</td>
      <td valign="middle">Invalid<br>in 7,573.08 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">-@a..com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be two . characters in a row in the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,420.49 ns</td>
      <td valign="middle">Invalid<br>in 9,188.13 ns</td>
      <td valign="middle">Invalid<br>in 35,778.07 ns</td>
      <td valign="middle">Invalid<br>in 11,117.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">invalid@about.museum-<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The top level domain cannot end with a - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 137.93 ns</td>
      <td valign="middle">Invalid<br>in 17,406.3 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,827.22 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 20,480.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@...........com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be more than one . character in a row</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,027.99 ns</td>
      <td valign="middle">Invalid<br>in 17,658.13 ns</td>
      <td valign="middle">Invalid<br>in 49,356.5 ns</td>
      <td valign="middle">Invalid<br>in 27,870.73 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@x2345678901234567890123456789<br>01234567890123456789012345678901234.test<br>.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A single domain part cannot have more than 63 characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 11,534.44 ns</td>
      <td valign="middle">Invalid<br>in 59,181.42 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 15,949.43 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 19,592.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">abc@def@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There can only be a single @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,143.68 ns</td>
      <td valign="middle">Invalid<br>in 3,104.65 ns</td>
      <td valign="middle">Invalid<br>in 20,671.36 ns</td>
      <td valign="middle">Invalid<br>in 10,585.42 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">abc\\@def@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The first @ character should be escaped</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,350.62 ns</td>
      <td valign="middle">Invalid<br>in 3,086.98 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,452.56 ns</td>
      <td valign="middle">Invalid<br>in 8,036.77 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">abc\@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Since the only @ is escaped, this is missing an @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,301.01 ns</td>
      <td valign="middle">Invalid<br>in 4,046.77 ns</td>
      <td valign="middle">Invalid<br>in 38,242.08 ns</td>
      <td valign="middle">Invalid<br>in 13,241.25 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot be empty</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,171.83 ns</td>
      <td valign="middle">Invalid<br>in 708.68 ns</td>
      <td valign="middle">Invalid<br>in 42,633.95 ns</td>
      <td valign="middle">Invalid<br>in 9,736.94 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">doug@<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot be empty</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,639.88 ns</td>
      <td valign="middle">Invalid<br>in 968.03 ns</td>
      <td valign="middle">Invalid<br>in 33,757.03 ns</td>
      <td valign="middle">Invalid<br>in 12,963.92 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">.dot@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot start with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 153.63 ns</td>
      <td valign="middle">Invalid<br>in 1,646.32 ns</td>
      <td valign="middle">Invalid<br>in 27,384.15 ns</td>
      <td valign="middle">Invalid<br>in 7,422.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">dot.@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot end with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,929.86 ns</td>
      <td valign="middle">Invalid<br>in 2,732.07 ns</td>
      <td valign="middle">Invalid<br>in 25,336.82 ns</td>
      <td valign="middle">Invalid<br>in 10,621.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">two..dot@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot have two . characters in a row</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,153.84 ns</td>
      <td valign="middle">Invalid<br>in 6,830.3 ns</td>
      <td valign="middle">Invalid<br>in 51,781.27 ns</td>
      <td valign="middle">Invalid<br>in 15,414.83 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Doug "Ace" L."@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Quotes must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,503.02 ns</td>
      <td valign="middle">Invalid<br>in 3,442.14 ns</td>
      <td valign="middle">Invalid<br>in 39,848.68 ns</td>
      <td valign="middle">Invalid<br>in 15,215.31 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Doug\ \"Ace\"\ L\.@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,248.13 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 8,787.52 ns</td>
      <td valign="middle">Invalid<br>in 28,085.18 ns</td>
      <td valign="middle">Invalid<br>in 15,902.59 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">hello world@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,731 ns</td>
      <td valign="middle">Invalid<br>in 3,390.26 ns</td>
      <td valign="middle">Invalid<br>in 25,055.1 ns</td>
      <td valign="middle">Invalid<br>in 15,838.82 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">gatsby@f.sc.ot.t.f.i.tzg.era.l.d.<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot end with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 125.66 ns</td>
      <td valign="middle">Invalid<br>in 74.02 ns</td>
      <td valign="middle">Invalid<br>in 41,903.96 ns</td>
      <td valign="middle">Invalid<br>in 32,611.73 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">.@<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot be a single . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 345.78 ns</td>
      <td valign="middle">Invalid<br>in 517.11 ns</td>
      <td valign="middle">Invalid<br>in 39,253.1 ns</td>
      <td valign="middle">Invalid<br>in 6,607.19 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">@bar.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot be empty</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 627.66 ns</td>
      <td valign="middle">Invalid<br>in 585 ns</td>
      <td valign="middle">Invalid<br>in 22,742.56 ns</td>
      <td valign="middle">Invalid<br>in 5,080.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">@@bar.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be more than one @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 398.34 ns</td>
      <td valign="middle">Invalid<br>in 3,580.66 ns</td>
      <td valign="middle">Invalid<br>in 29,164.01 ns</td>
      <td valign="middle">Invalid<br>in 9,004.57 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">aaa@.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot start with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 999.06 ns</td>
      <td valign="middle">Invalid<br>in 4,320.05 ns</td>
      <td valign="middle">Invalid<br>in 20,625.14 ns</td>
      <td valign="middle">Invalid<br>in 8,823.07 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">aaa@.123<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot start with a . character and the tld cannot be all numeric</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 992.85 ns</td>
      <td valign="middle">Invalid<br>in 3,921.92 ns</td>
      <td valign="middle">Invalid<br>in 20,645.55 ns</td>
      <td valign="middle">Invalid<br>in 11,249.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a@bar.com.<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot end with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 130.13 ns</td>
      <td valign="middle">Invalid<br>in 71.69 ns</td>
      <td valign="middle">Invalid<br>in 25,721.76 ns</td>
      <td valign="middle">Invalid<br>in 13,837.67 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(foo(bar)@iamcal.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The comment must be closed</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,719.54 ns</td>
      <td valign="middle">Invalid<br>in 3,664.73 ns</td>
      <td valign="middle">Invalid<br>in 22,906.59 ns</td>
      <td valign="middle">Invalid<br>in 17,427.54 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(foo)bar)@iamcal.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There is an extra closing parenthesis</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,239.82 ns</td>
      <td valign="middle">Invalid<br>in 3,927.01 ns</td>
      <td valign="middle">Invalid<br>in 30,547.26 ns</td>
      <td valign="middle">Invalid<br>in 17,037.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(foo\)@iamcal.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The closing parenthesis is escaped so there is no actual closing parenthesis</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,479.48 ns</td>
      <td valign="middle">Invalid<br>in 3,375.47 ns</td>
      <td valign="middle">Invalid<br>in 26,960.98 ns</td>
      <td valign="middle">Invalid<br>in 46,161.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(1234567890123456789012345678901234<br>5678901234567890)last@(12345678901234567<br>8901234567890123456789012345678901234567<br>8901234567890123456789012345678901234567<br>8901234567890123456789012345678901234567<br>8901234567890123456789012345678901234567<br>8901234567890123456789012345678901234567<br>890123456789012345678901234567890)test.o<br>rg<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The comment in the local-part must be . separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 167.59 ns</td>
      <td valign="middle">Invalid<br>in 19,050.86 ns</td>
      <td valign="middle">Invalid<br>in 51,944.62 ns</td>
      <td valign="middle">Invalid<br>in 134,817.53 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a(a(b(c)d(e(f))g)(h(i)j)@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Nested comments are ok, but this is missing a closing parenthesis</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,583.31 ns</td>
      <td valign="middle">Invalid<br>in 2,155.69 ns</td>
      <td valign="middle">Invalid<br>in 22,900.68 ns</td>
      <td valign="middle">Invalid<br>in 10,571.47 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Doug\ \"Ace\"\ Lovell@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,068.67 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 11,613.76 ns</td>
      <td valign="middle">Invalid<br>in 25,059.63 ns</td>
      <td valign="middle">Invalid<br>in 11,775.7 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test.test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There must be an @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,042.4 ns</td>
      <td valign="middle">Invalid<br>in 1,435.79 ns</td>
      <td valign="middle">Invalid<br>in 36,581.84 ns</td>
      <td valign="middle">Invalid<br>in 25,654.33 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test.@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot end with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,971.17 ns</td>
      <td valign="middle">Invalid<br>in 3,010.99 ns</td>
      <td valign="middle">Invalid<br>in 22,406.2 ns</td>
      <td valign="middle">Invalid<br>in 13,466.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test..test@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot have two . characters in a row</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,458.27 ns</td>
      <td valign="middle">Invalid<br>in 7,400.71 ns</td>
      <td valign="middle">Invalid<br>in 33,207.37 ns</td>
      <td valign="middle">Invalid<br>in 15,155.35 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">.test@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot start with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 160.21 ns</td>
      <td valign="middle">Invalid<br>in 2,975.39 ns</td>
      <td valign="middle">Invalid<br>in 39,606.18 ns</td>
      <td valign="middle">Invalid<br>in 15,072.33 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@test@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be more than one @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,787.09 ns</td>
      <td valign="middle">Invalid<br>in 4,220.29 ns</td>
      <td valign="middle">Invalid<br>in 35,336.62 ns</td>
      <td valign="middle">Invalid<br>in 23,996.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">There cannot be more than one @ character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,030.06 ns</td>
      <td valign="middle">Invalid<br>in 3,186.63 ns</td>
      <td valign="middle">Invalid<br>in 21,244.73 ns</td>
      <td valign="middle">Invalid<br>in 10,029.45 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">-- test --@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,676.14 ns</td>
      <td valign="middle">Invalid<br>in 2,561.59 ns</td>
      <td valign="middle">Invalid<br>in 20,970 ns</td>
      <td valign="middle">Invalid<br>in 14,774.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@te st.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,147.04 ns</td>
      <td valign="middle">Invalid<br>in 3,909.38 ns</td>
      <td valign="middle">Invalid<br>in 38,645.15 ns</td>
      <td valign="middle">Invalid<br>in 15,188.11 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">[test]@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The characters [ and ] are not allowed unquoted</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,077.89 ns</td>
      <td valign="middle">Invalid<br>in 5,499.27 ns</td>
      <td valign="middle">Invalid<br>in 42,631.45 ns</td>
      <td valign="middle">Invalid<br>in 11,308.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test"test"@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Quotes must be in pairs</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,699.63 ns</td>
      <td valign="middle">Invalid<br>in 7,329.02 ns</td>
      <td valign="middle">Invalid<br>in 41,209.7 ns</td>
      <td valign="middle">Invalid<br>in 15,540.6 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">()[]\;:,&gt;&lt;@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part contains invalid unquoted characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,547.03 ns</td>
      <td valign="middle">Invalid<br>in 6,027.51 ns</td>
      <td valign="middle">Invalid<br>in 32,121.7 ns</td>
      <td valign="middle">Invalid<br>in 25,518.64 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@.<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot be a single . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 193.2 ns</td>
      <td valign="middle">Invalid<br>in 89.33 ns</td>
      <td valign="middle">Invalid<br>in 27,033.54 ns</td>
      <td valign="middle">Invalid<br>in 11,037.92 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@example.<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot end with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 146.7 ns</td>
      <td valign="middle">Invalid<br>in 85.45 ns</td>
      <td valign="middle">Invalid<br>in 36,519.26 ns</td>
      <td valign="middle">Invalid<br>in 15,092.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot start with a . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,599.62 ns</td>
      <td valign="middle">Invalid<br>in 4,936.48 ns</td>
      <td valign="middle">Invalid<br>in 28,790.23 ns</td>
      <td valign="middle">Invalid<br>in 15,675.72 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@12345678901234567890123456789012345<br>6789012345678901234567890123456789012345<br>6789012345678901234567890123456789012345<br>6789012345678901234567890123456789012345<br>6789012345678901234567890123456789012345<br>6789012345678901234567890123456789012345<br>67890123456789012.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Each domain part can only be 63 characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 35,762.28 ns</td>
      <td valign="middle">Invalid<br>in 55,549.41 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 70,387.23 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 37,433.21 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">.@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The local-part cannot be a single . character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 205.76 ns</td>
      <td valign="middle">Invalid<br>in 3,032.85 ns</td>
      <td valign="middle">Invalid<br>in 46,059.51 ns</td>
      <td valign="middle">Invalid<br>in 11,523.71 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Ima Fool@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Whitespace must be separated by . characters</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,771.69 ns</td>
      <td valign="middle">Invalid<br>in 4,261.14 ns</td>
      <td valign="middle">Invalid<br>in 34,458.85 ns</td>
      <td valign="middle">Invalid<br>in 22,777.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first\\"last"@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A " character within quotes must be escaped</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,812.22 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,701.98 ns</td>
      <td valign="middle">Invalid<br>in 20,116.36 ns</td>
      <td valign="middle">Invalid<br>in 10,634.25 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">foo@[\1.2.3.4]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The \ character is not allowed within the domain literal</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,168.71 ns</td>
      <td valign="middle">Invalid<br>in 5,467.78 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 1,754.29 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 10,995.45 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first\last@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The \ character is not escaping anything</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,646.8 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 7,542.31 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,624.22 ns</td>
      <td valign="middle">Invalid<br>in 9,938.23 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(abc("def".ghi).mno)middle(abc("def<br>".ghi).mno).last@(abc("def".ghi).mno)exa<br>mple(abc("def".ghi).mno).(abc("def".ghi)<br>.mno)com(abc("def".ghi).mno)<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The first comment is not dot separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,113.16 ns</td>
      <td valign="middle">Invalid<br>in 21,910.37 ns</td>
      <td valign="middle">Invalid<br>in 44,399.24 ns</td>
      <td valign="middle">Invalid<br>in 28,700.89 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(middle)last@test.org<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Comments in the middle of the local-part must be dot separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,730.13 ns</td>
      <td valign="middle">Invalid<br>in 6,134.72 ns</td>
      <td valign="middle">Invalid<br>in 39,664.61 ns</td>
      <td valign="middle">Invalid<br>in 34,783.98 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Unicode NULL ␀"@char.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The unicode null character must be escaped within quotes</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,805.66 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 15,294.54 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 3,735.57 ns</td>
      <td valign="middle">Invalid<br>in 19,961.83 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Unicode NULL \␀@char.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The escaped unicode null character is only allowed within quotes</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,153.84 ns</td>
      <td valign="middle">Invalid<br>in 5,827.53 ns</td>
      <td valign="middle">Invalid<br>in 21,943.72 ns</td>
      <td valign="middle">Invalid<br>in 26,734.33 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test"test@test.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Quotes must be dot separated</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 1,296.71 ns</td>
      <td valign="middle">Invalid<br>in 3,552.64 ns</td>
      <td valign="middle">Invalid<br>in 92,596.2 ns</td>
      <td valign="middle">Invalid<br>in 19,135.43 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">()@test.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Empty comment is not allowed</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 4,181.92 ns</td>
      <td valign="middle">Invalid<br>in 4,842.05 ns</td>
      <td valign="middle">Invalid<br>in 32,670.64 ns</td>
      <td valign="middle">Invalid<br>in 14,405.87 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@really.long.topleveldomainisnotallo<br>wedunfortunatelyforpeoplewholikereallylo<br>ngtopleveldomainnames<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The top level domain can only be 63 characters long</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 12,742.47 ns</td>
      <td valign="middle">Invalid<br>in 36,608.69 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 14,544.1 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 18,118.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@really.long.domainpartisnotallowedu<br>nfortunatelyforpeoplewholikereallylongdo<br>mainnameparts.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Each domain part can only be 63 characters long</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 9,111.92 ns</td>
      <td valign="middle">Invalid<br>in 23,723.12 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 11,009.81 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 17,850.32 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">invalid@[1]<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain literal is not a valid IPv4 or IPv6 address</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,241.58 ns</td>
      <td valign="middle">Invalid<br>in 7,074 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 2,043.59 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 16,965.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">ä📧@-foo<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A dotless domain cannot start with the - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 7,940.85 ns</td>
      <td valign="middle">Invalid<br>in 11,792.6 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,256.26 ns</td>
      <td valign="middle">Invalid<br>in 10,344.86 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">ä📧@foo-<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A dotless domain cannot end with the - character</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 190.27 ns</td>
      <td valign="middle">Invalid<br>in 12,341.77 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 1,665.29 ns</td>
      <td valign="middle">Invalid<br>in 6,536.58 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(comment(inner@comment.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">Comments must have closing parenthesis</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,066.49 ns</td>
      <td valign="middle">Invalid<br>in 5,528.99 ns</td>
      <td valign="middle">Invalid<br>in 20,075.99 ns</td>
      <td valign="middle">Invalid<br>in 49,957.11 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith &lt;email@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A quoted identifier address must have a closing angled bracket</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,715.65 ns</td>
      <td valign="middle">Invalid<br>in 6,867.86 ns</td>
      <td valign="middle">Invalid<br>in 29,946.7 ns</td>
      <td valign="middle">Invalid<br>in 75,219.83 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith email@example.com<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The address part must be in angled brackets</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 3,070.73 ns</td>
      <td valign="middle">Invalid<br>in 4,481.38 ns</td>
      <td valign="middle">Invalid<br>in 22,222.88 ns</td>
      <td valign="middle">Invalid<br>in 28,278.43 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith &lt;email@example.com-&gt;<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The address part must be valid (cannot end with a dash)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,257.34 ns</td>
      <td valign="middle">Invalid<br>in 4,720.71 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 5,256.86 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 16,670.14 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith &lt;email@-example.com-&gt;<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The address part must be valid (domain cannot start with a dash)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,436.71 ns</td>
      <td valign="middle">Invalid<br>in 4,316.89 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 4,142.79 ns</td>
      <td style="background-color:#815355" valign="middle">Valid<br>in 18,060.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith &lt;email&gt;<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The address part must be valid (missing @ character and domain)</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 6,259.42 ns</td>
      <td valign="middle">Invalid<br>in 4,340.79 ns</td>
      <td valign="middle">Invalid<br>in 54,148.69 ns</td>
      <td valign="middle">Invalid<br>in 51,485.38 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">ABC.DEF@GHI. (MNO)<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">The domain cannot end with the . character even with an ending comment</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 5,616.77 ns</td>
      <td valign="middle">Invalid<br>in 2,945.22 ns</td>
      <td valign="middle">Invalid<br>in 30,359.21 ns</td>
      <td valign="middle">Invalid<br>in 24,961.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@test.com(comment<br></th>
      <td valign="middle">Invalid<br><small class="text-muted">A comment at the end of the domain must have closing parenthesis</small></td>
      <td style="background-color:#56666B" valign="middle">Invalid<br>in 2,511.35 ns</td>
      <td valign="middle">Invalid<br>in 9,389.9 ns</td>
      <td valign="middle">Invalid<br>in 20,484.27 ns</td>
      <td valign="middle">Invalid<br>in 165,800.42 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">" "@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,561.98 ns</td>
      <td valign="middle">Valid<br>in 7,166.08 ns</td>
      <td valign="middle">Valid<br>in 3,264.57 ns</td>
      <td valign="middle">Valid<br>in 11,362.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"john..doe"@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,288.08 ns</td>
      <td valign="middle">Valid<br>in 11,524.86 ns</td>
      <td valign="middle">Valid<br>in 3,738.55 ns</td>
      <td valign="middle">Valid<br>in 18,855.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"email"@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,521.48 ns</td>
      <td valign="middle">Valid<br>in 11,041.68 ns</td>
      <td valign="middle">Valid<br>in 4,199.5 ns</td>
      <td valign="middle">Valid<br>in 24,389.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first@last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,286.32 ns</td>
      <td valign="middle">Valid<br>in 14,209.35 ns</td>
      <td valign="middle">Valid<br>in 6,084.05 ns</td>
      <td valign="middle">Valid<br>in 17,512.41 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">very.unusual."@".unusual.com@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,370.64 ns</td>
      <td valign="middle">Valid<br>in 12,366.41 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 22,250.87 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 19,739.64 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first\"last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,833.31 ns</td>
      <td valign="middle">Valid<br>in 6,376.55 ns</td>
      <td valign="middle">Valid<br>in 2,967.24 ns</td>
      <td valign="middle">Valid<br>in 15,155.02 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">much."more\ unusual"@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,535.25 ns</td>
      <td valign="middle">Valid<br>in 10,696.61 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 31,015.66 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 16,324.71 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first\\last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,799.58 ns</td>
      <td valign="middle">Valid<br>in 6,177.36 ns</td>
      <td valign="middle">Valid<br>in 2,283.84 ns</td>
      <td valign="middle">Valid<br>in 13,654.23 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Abc\@def"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,067.7 ns</td>
      <td valign="middle">Valid<br>in 8,351.98 ns</td>
      <td valign="middle">Valid<br>in 2,173.13 ns</td>
      <td valign="middle">Valid<br>in 20,871.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Fred\ Bloggs"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 6,473.04 ns</td>
      <td valign="middle">Valid<br>in 10,486.57 ns</td>
      <td valign="middle">Valid<br>in 3,724.47 ns</td>
      <td valign="middle">Valid<br>in 23,055.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Joe.\\Blow"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,503.87 ns</td>
      <td valign="middle">Valid<br>in 6,708.93 ns</td>
      <td valign="middle">Valid<br>in 2,540.2 ns</td>
      <td valign="middle">Valid<br>in 15,487.28 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Abc@def"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 16,109.81 ns</td>
      <td valign="middle">Valid<br>in 12,973.94 ns</td>
      <td valign="middle">Valid<br>in 3,962.45 ns</td>
      <td valign="middle">Valid<br>in 31,331.37 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Fred Bloggs"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,841.43 ns</td>
      <td valign="middle">Valid<br>in 6,284.22 ns</td>
      <td valign="middle">Valid<br>in 2,338.23 ns</td>
      <td valign="middle">Valid<br>in 16,467.74 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first\last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,826.58 ns</td>
      <td valign="middle">Valid<br>in 6,176.51 ns</td>
      <td valign="middle">Valid<br>in 3,271.01 ns</td>
      <td valign="middle">Valid<br>in 24,366.73 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Doug \"Ace\" L."@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,099.41 ns</td>
      <td valign="middle">Valid<br>in 8,134.87 ns</td>
      <td valign="middle">Valid<br>in 2,402.92 ns</td>
      <td valign="middle">Valid<br>in 14,399.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"[[ test ]]"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,986.71 ns</td>
      <td valign="middle">Valid<br>in 5,834.58 ns</td>
      <td valign="middle">Valid<br>in 3,001.12 ns</td>
      <td valign="middle">Valid<br>in 14,267.35 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test.test"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,116.29 ns</td>
      <td valign="middle">Valid<br>in 5,612.63 ns</td>
      <td valign="middle">Valid<br>in 2,190.9 ns</td>
      <td valign="middle">Valid<br>in 12,301.87 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test."test"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,855.15 ns</td>
      <td valign="middle">Valid<br>in 11,673.31 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 29,779.77 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 15,505.69 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test@test"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,792.74 ns</td>
      <td valign="middle">Valid<br>in 6,266.78 ns</td>
      <td valign="middle">Valid<br>in 3,264.36 ns</td>
      <td valign="middle">Valid<br>in 14,227.75 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test\test"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,200.77 ns</td>
      <td valign="middle">Valid<br>in 5,914.4 ns</td>
      <td valign="middle">Valid<br>in 2,729.41 ns</td>
      <td valign="middle">Valid<br>in 19,584.94 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first"."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,421.58 ns</td>
      <td valign="middle">Valid<br>in 10,743.72 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 35,304.02 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 15,320.59 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first".middle."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,630.16 ns</td>
      <td valign="middle">Valid<br>in 11,244.91 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 32,174.36 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 16,663.49 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first".last@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,913.17 ns</td>
      <td valign="middle">Valid<br>in 10,675.11 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 30,281.32 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 16,384.67 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,293.13 ns</td>
      <td valign="middle">Valid<br>in 11,402.38 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 32,067.72 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 21,617.12 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first"."middle"."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,707.32 ns</td>
      <td valign="middle">Valid<br>in 8,171.08 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 27,954.03 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,091.98 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first.middle"."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,623.68 ns</td>
      <td valign="middle">Valid<br>in 9,159.2 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 32,096.68 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 25,305.13 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first.middle.last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,647.88 ns</td>
      <td valign="middle">Valid<br>in 12,659.43 ns</td>
      <td valign="middle">Valid<br>in 5,503.47 ns</td>
      <td valign="middle">Valid<br>in 24,133.86 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first..last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,761.07 ns</td>
      <td valign="middle">Valid<br>in 13,895.27 ns</td>
      <td valign="middle">Valid<br>in 2,866.88 ns</td>
      <td valign="middle">Valid<br>in 26,323.61 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Unicode NULL \␀"@char.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,072.15 ns</td>
      <td valign="middle">Valid<br>in 15,312.63 ns</td>
      <td valign="middle">Valid<br>in 3,293.1 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 24,764.36 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test\\blah"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,521.63 ns</td>
      <td valign="middle">Valid<br>in 16,836.22 ns</td>
      <td valign="middle">Valid<br>in 2,662.77 ns</td>
      <td valign="middle">Valid<br>in 24,375.58 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test\blah"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,383.81 ns</td>
      <td valign="middle">Valid<br>in 11,374.14 ns</td>
      <td valign="middle">Valid<br>in 3,495.52 ns</td>
      <td valign="middle">Valid<br>in 26,831.72 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test\"blah"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,851.77 ns</td>
      <td valign="middle">Valid<br>in 10,323.5 ns</td>
      <td valign="middle">Valid<br>in 3,878.07 ns</td>
      <td valign="middle">Valid<br>in 26,268.4 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"first\\\"last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,397.39 ns</td>
      <td valign="middle">Valid<br>in 18,673.24 ns</td>
      <td valign="middle">Valid<br>in 3,451.69 ns</td>
      <td valign="middle">Valid<br>in 31,855.22 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first."mid\dle"."last"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,510.63 ns</td>
      <td valign="middle">Valid<br>in 24,051.49 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 42,490.14 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 25,174.58 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"Test \"Fail\" Ing"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,165.31 ns</td>
      <td valign="middle">Valid<br>in 16,291.07 ns</td>
      <td valign="middle">Valid<br>in 4,348.5 ns</td>
      <td valign="middle">Valid<br>in 35,658.65 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"test
 blah"@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,148.73 ns</td>
      <td valign="middle">Valid<br>in 17,991.49 ns</td>
      <td valign="middle">Valid<br>in 5,291.74 ns</td>
      <td valign="middle">Valid<br>in 28,332.35 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last @test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,413.02 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 6,958.49 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 36,558.89 ns</td>
      <td valign="middle">Valid<br>in 35,430.33 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last  @test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,381.15 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 10,402.04 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 36,825.29 ns</td>
      <td valign="middle">Valid<br>in 44,005.64 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first .last  @test .org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,839.05 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,362.51 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 38,358.74 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 43,464.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">jdoe@machine(comment).  example<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,654.28 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 6,624.24 ns</td>
      <td valign="middle">Valid<br>in 5,204.09 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 36,409.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">very."(),:;&lt;&gt;[]".VERY."very@\\ ”ve<br>ry".unusual@strange.example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 34,896.38 ns</td>
      <td valign="middle">Valid<br>in 25,040.51 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 46,138.67 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 14,657.83 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first."".last@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,713.07 ns</td>
      <td valign="middle">Valid<br>in 17,162.44 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 41,273.26 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 15,671.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">""@test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted"></small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,613.08 ns</td>
      <td valign="middle">Valid<br>in 8,507.08 ns</td>
      <td valign="middle">Valid<br>in 2,417.94 ns</td>
      <td valign="middle">Valid<br>in 18,472.36 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">simple@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A simple valid email address</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 21,895.18 ns</td>
      <td valign="middle">Valid<br>in 8,365.15 ns</td>
      <td valign="middle">Valid<br>in 6,654.99 ns</td>
      <td valign="middle">Valid<br>in 16,455.43 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">very.common@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A simple address with a dot separator</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,326.9 ns</td>
      <td valign="middle">Valid<br>in 18,798.78 ns</td>
      <td valign="middle">Valid<br>in 7,709.95 ns</td>
      <td valign="middle">Valid<br>in 29,099.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">very.common@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A simple address ending in .org</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,856.13 ns</td>
      <td valign="middle">Valid<br>in 20,971.65 ns</td>
      <td valign="middle">Valid<br>in 7,889.54 ns</td>
      <td valign="middle">Valid<br>in 30,302.05 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">disposable.style.email.with+symbol@examp<br>le.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The + symbol is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,387.97 ns</td>
      <td valign="middle">Valid<br>in 20,063.77 ns</td>
      <td valign="middle">Valid<br>in 14,869.3 ns</td>
      <td valign="middle">Valid<br>in 66,068.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">other.email-with-hyphen@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The hyphen symbol is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,017.38 ns</td>
      <td valign="middle">Valid<br>in 24,116.66 ns</td>
      <td valign="middle">Valid<br>in 8,819.47 ns</td>
      <td valign="middle">Valid<br>in 50,937.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">fully-qualified-domain@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The hyphen symbol is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 18,169.78 ns</td>
      <td valign="middle">Valid<br>in 20,694.88 ns</td>
      <td valign="middle">Valid<br>in 11,059.97 ns</td>
      <td valign="middle">Valid<br>in 44,313.97 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">user.name+tag+sorting@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">Tag-style emails with + are valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,603.81 ns</td>
      <td valign="middle">Valid<br>in 18,225.02 ns</td>
      <td valign="middle">Valid<br>in 10,173.18 ns</td>
      <td valign="middle">Valid<br>in 43,307.68 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">x@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">Single character local-part is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,903.97 ns</td>
      <td valign="middle">Valid<br>in 9,096.64 ns</td>
      <td valign="middle">Valid<br>in 5,130.16 ns</td>
      <td valign="middle">Valid<br>in 14,650.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">example-indeed@strange-example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A hyphen is allowed in the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,219.03 ns</td>
      <td valign="middle">Valid<br>in 23,103.91 ns</td>
      <td valign="middle">Valid<br>in 8,598.19 ns</td>
      <td valign="middle">Valid<br>in 42,126.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test/test@test.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A forward slash is allowed in the local-part</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,597.17 ns</td>
      <td valign="middle">Valid<br>in 15,752.08 ns</td>
      <td valign="middle">Valid<br>in 6,619.38 ns</td>
      <td valign="middle">Valid<br>in 23,707.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">admin@mailserver1<br></th>
      <td valign="middle">Valid<br><small class="text-muted">Top-level domain is not required</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,671.04 ns</td>
      <td valign="middle">Valid<br>in 17,590.23 ns</td>
      <td valign="middle">Valid<br>in 5,739.13 ns</td>
      <td valign="middle">Valid<br>in 17,320.89 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">example@s.example<br></th>
      <td valign="middle">Valid<br><small class="text-muted">.example is a valid TLD</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,351.71 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,878.81 ns</td>
      <td valign="middle">Valid<br>in 4,959.83 ns</td>
      <td valign="middle">Valid<br>in 23,994.94 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">mailhost!username@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The ! character is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,693.47 ns</td>
      <td valign="middle">Valid<br>in 22,861.16 ns</td>
      <td valign="middle">Valid<br>in 7,349.36 ns</td>
      <td valign="middle">Valid<br>in 35,284.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">user%example.com@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The % character is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,517.19 ns</td>
      <td valign="middle">Valid<br>in 23,132.28 ns</td>
      <td valign="middle">Valid<br>in 9,551.61 ns</td>
      <td valign="middle">Valid<br>in 35,531.32 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">user-@example.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The local-part can end in a hyphen</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,560.86 ns</td>
      <td valign="middle">Valid<br>in 9,641.27 ns</td>
      <td valign="middle">Valid<br>in 4,202.58 ns</td>
      <td valign="middle">Valid<br>in 15,371.42 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A simple valid email address</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,464.74 ns</td>
      <td valign="middle">Valid<br>in 8,173.63 ns</td>
      <td valign="middle">Valid<br>in 4,787.86 ns</td>
      <td valign="middle">Valid<br>in 21,367.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">firstname.lastname@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A simple address with a dot separator</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,482.93 ns</td>
      <td valign="middle">Valid<br>in 24,410.51 ns</td>
      <td valign="middle">Valid<br>in 8,793.51 ns</td>
      <td valign="middle">Valid<br>in 39,711.56 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@subdomain.example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">An email address with a subdomain is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 16,446.65 ns</td>
      <td valign="middle">Valid<br>in 15,983.01 ns</td>
      <td valign="middle">Valid<br>in 5,874.62 ns</td>
      <td valign="middle">Valid<br>in 20,003.59 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">firstname+lastname@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The + character is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,247.83 ns</td>
      <td valign="middle">Valid<br>in 16,060.41 ns</td>
      <td valign="middle">Valid<br>in 10,351.04 ns</td>
      <td valign="middle">Valid<br>in 38,175.47 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">1234567890@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">An all-numeric local-part is valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,235.07 ns</td>
      <td valign="middle">Valid<br>in 16,158.65 ns</td>
      <td valign="middle">Valid<br>in 6,969.06 ns</td>
      <td valign="middle">Valid<br>in 25,901.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example-one.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A hyphen is allowed in specific parts of the domain</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,978.27 ns</td>
      <td valign="middle">Valid<br>in 12,668.79 ns</td>
      <td valign="middle">Valid<br>in 6,932.16 ns</td>
      <td valign="middle">Valid<br>in 20,091.7 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">_______@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The underscore character is allowed in the local-part</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,520.28 ns</td>
      <td valign="middle">Valid<br>in 42,621.92 ns</td>
      <td valign="middle">Valid<br>in 4,749.85 ns</td>
      <td valign="middle">Valid<br>in 26,524.32 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example.name<br></th>
      <td valign="middle">Valid<br><small class="text-muted">.name is a valid TLD</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,727.3 ns</td>
      <td valign="middle">Valid<br>in 13,491.22 ns</td>
      <td valign="middle">Valid<br>in 4,579.22 ns</td>
      <td valign="middle">Valid<br>in 20,671.84 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example.museum<br></th>
      <td valign="middle">Valid<br><small class="text-muted">.museum is a valid TLD</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 18,693.46 ns</td>
      <td valign="middle">Valid<br>in 14,197.04 ns</td>
      <td valign="middle">Valid<br>in 4,714.66 ns</td>
      <td valign="middle">Valid<br>in 27,131.71 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example.co.jp<br></th>
      <td valign="middle">Valid<br><small class="text-muted">.co.jp is a valid TLD</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,126.31 ns</td>
      <td valign="middle">Valid<br>in 19,308.74 ns</td>
      <td valign="middle">Valid<br>in 5,504.68 ns</td>
      <td valign="middle">Valid<br>in 20,170.31 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">firstname-lastname@example.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A hyphen is allowed in the local-part</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,127.36 ns</td>
      <td valign="middle">Valid<br>in 20,054.36 ns</td>
      <td valign="middle">Valid<br>in 9,187.09 ns</td>
      <td valign="middle">Valid<br>in 44,562.95 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">x@x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x23456789.x23456789.x2345678<br>9.x23456789.x2<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A domain under 256 characters is allowed</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 85,580.67 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 79,824.6 ns</td>
      <td valign="middle">Valid<br>in 37,800.22 ns</td>
      <td valign="middle">Valid<br>in 28,984.78 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">1234567890123456789012345678901234567890<br>123456789012345678@123456789012345678901<br>23456789012345678901234567890123456789.1<br>2345678901234567890123456789012345678901<br>234567890123456789.123456789012345678901<br>2345678901234567890123456789012345678901<br>23.test.org<br></th>
      <td valign="middle">Valid<br><small class="text-muted">A local-part under 65 characters is allowed</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 81,373.51 ns</td>
      <td valign="middle">Valid<br>in 76,734.46 ns</td>
      <td valign="middle">Valid<br>in 38,660.81 ns</td>
      <td valign="middle">Valid<br>in 88,060.57 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@3com.com<br></th>
      <td valign="middle">Valid<br><small class="text-muted">The domain can start with a number</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,673.54 ns</td>
      <td valign="middle">Valid<br>in 7,616.62 ns</td>
      <td valign="middle">Valid<br>in 5,867.44 ns</td>
      <td valign="middle">Valid<br>in 28,305.95 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@123.test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,168.39 ns</td>
      <td valign="middle">Valid<br>in 11,850.63 ns</td>
      <td valign="middle">Valid<br>in 6,540.56 ns</td>
      <td valign="middle">Valid<br>in 28,836.16 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@x2345678901234567890123456789<br>0123456789012345678901234567890123.test.<br>org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 33,688.22 ns</td>
      <td valign="middle">Valid<br>in 28,154.54 ns</td>
      <td valign="middle">Valid<br>in 16,338.23 ns</td>
      <td valign="middle">Valid<br>in 32,925.76 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">1234567890123456789012345678901234567890<br>123456789012345678901234@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,254.33 ns</td>
      <td valign="middle">Valid<br>in 39,110.09 ns</td>
      <td valign="middle">Valid<br>in 27,778.82 ns</td>
      <td valign="middle">Valid<br>in 114,039.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@[123.123.123.123]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 4,571.33 ns</td>
      <td valign="middle">Valid<br>in 10,984.67 ns</td>
      <td valign="middle">Valid<br>in 7,884.49 ns</td>
      <td valign="middle">Valid<br>in 25,894.36 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[12.34.56.78]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,515.26 ns</td>
      <td valign="middle">Valid<br>in 15,969.45 ns</td>
      <td valign="middle">Valid<br>in 6,868.39 ns</td>
      <td valign="middle">Valid<br>in 25,858.21 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">user+mailbox@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,951.99 ns</td>
      <td valign="middle">Valid<br>in 12,177.65 ns</td>
      <td valign="middle">Valid<br>in 6,974.5 ns</td>
      <td valign="middle">Valid<br>in 27,564.23 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">customer/department=shipping@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 35,768.95 ns</td>
      <td valign="middle">Valid<br>in 32,529.32 ns</td>
      <td valign="middle">Valid<br>in 6,836.32 ns</td>
      <td valign="middle">Valid<br>in 54,540.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">$A12345@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,346.93 ns</td>
      <td valign="middle">Valid<br>in 11,738.12 ns</td>
      <td valign="middle">Valid<br>in 2,984.9 ns</td>
      <td valign="middle">Valid<br>in 27,739.84 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">!def!xyz%abc@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 29,701.13 ns</td>
      <td valign="middle">Valid<br>in 17,505.57 ns</td>
      <td valign="middle">Valid<br>in 8,759.08 ns</td>
      <td valign="middle">Valid<br>in 45,331.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">_somename@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 23,992.04 ns</td>
      <td valign="middle">Valid<br>in 22,523.65 ns</td>
      <td valign="middle">Valid<br>in 7,560 ns</td>
      <td valign="middle">Valid<br>in 37,290.15 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">dclo@us.ibm.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 25,957.1 ns</td>
      <td valign="middle">Valid<br>in 17,134.22 ns</td>
      <td valign="middle">Valid<br>in 6,069.65 ns</td>
      <td valign="middle">Valid<br>in 21,864.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@xn--example.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 30,294.16 ns</td>
      <td valign="middle">Valid<br>in 20,690.28 ns</td>
      <td valign="middle">Valid<br>in 4,787.53 ns</td>
      <td valign="middle">Valid<br>in 21,708.08 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:7777:8888]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 27,591.5 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 32,183.82 ns</td>
      <td valign="middle">Valid<br>in 10,438.21 ns</td>
      <td valign="middle">Valid<br>in 30,747.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333::4444:55<br>55:6666:7777]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 30,520.52 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 20,369.49 ns</td>
      <td valign="middle">Valid<br>in 6,327.56 ns</td>
      <td valign="middle">Valid<br>in 46,646.56 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::1111:2222:3333:4444:5<br>555:6666]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,347.46 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 13,655.79 ns</td>
      <td valign="middle">Valid<br>in 6,580.52 ns</td>
      <td valign="middle">Valid<br>in 30,434.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333::4444:55<br>55:6666]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 26,573.98 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 50,931.96 ns</td>
      <td valign="middle">Valid<br>in 4,732.98 ns</td>
      <td valign="middle">Valid<br>in 33,247.54 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 18,386.29 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 59,531.49 ns</td>
      <td valign="middle">Valid<br>in 6,220.05 ns</td>
      <td valign="middle">Valid<br>in 25,446.41 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::a2:a3:a4:b1:b2:b3:b4]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,955.49 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 8,211.41 ns</td>
      <td valign="middle">Valid<br>in 5,210.13 ns</td>
      <td valign="middle">Valid<br>in 26,044.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:a3:a4:b1:b2:b3::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,675.65 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 9,095.76 ns</td>
      <td valign="middle">Valid<br>in 4,092.96 ns</td>
      <td valign="middle">Valid<br>in 33,800.54 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,743.54 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 9,106.03 ns</td>
      <td valign="middle">Valid<br>in 3,760.69 ns</td>
      <td valign="middle">Valid<br>in 19,353.54 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::b4]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,151.15 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 5,843.97 ns</td>
      <td valign="middle">Valid<br>in 3,377.9 ns</td>
      <td valign="middle">Valid<br>in 30,453.13 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::b3:b4]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 6,401.26 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 7,010.08 ns</td>
      <td valign="middle">Valid<br>in 4,095.78 ns</td>
      <td valign="middle">Valid<br>in 24,387.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::b4]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,306.32 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 50,118.94 ns</td>
      <td valign="middle">Valid<br>in 3,470.3 ns</td>
      <td valign="middle">Valid<br>in 26,692.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,204.15 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 82,974.49 ns</td>
      <td valign="middle">Valid<br>in 4,312.32 ns</td>
      <td valign="middle">Valid<br>in 41,073.43 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,978.35 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 78,628.23 ns</td>
      <td valign="middle">Valid<br>in 5,586.79 ns</td>
      <td valign="middle">Valid<br>in 28,619.33 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:cdef::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,413.62 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 62,453.92 ns</td>
      <td valign="middle">Valid<br>in 5,041.31 ns</td>
      <td valign="middle">Valid<br>in 32,847.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:CDEF::]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 16,978.28 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 52,082.83 ns</td>
      <td valign="middle">Valid<br>in 5,171.89 ns</td>
      <td valign="middle">Valid<br>in 33,911.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::12.34.56.78]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,778.58 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 5,791.92 ns</td>
      <td valign="middle">Valid<br>in 4,085.55 ns</td>
      <td valign="middle">Valid<br>in 27,471.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333::4444:12<br>.34.56.78]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 19,669.51 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 71,303.61 ns</td>
      <td valign="middle">Valid<br>in 8,764.28 ns</td>
      <td valign="middle">Valid<br>in 38,413.67 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333:4444:555<br>5:6666:12.34.56.78]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,855.74 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 77,400.19 ns</td>
      <td valign="middle">Valid<br>in 7,860.16 ns</td>
      <td valign="middle">Valid<br>in 44,322.5 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:1111:2222:3333::4444:55<br>55:12.34.56.78]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 20,304.6 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 78,602.04 ns</td>
      <td valign="middle">Valid<br>in 2,410.23 ns</td>
      <td valign="middle">Valid<br>in 35,874.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">aaa@[123.123.123.123]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 6,188.83 ns</td>
      <td valign="middle">Valid<br>in 5,697.1 ns</td>
      <td valign="middle">Valid<br>in 1,867.63 ns</td>
      <td valign="middle">Valid<br>in 31,484.59 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::a3:a4:b1:ffff:11.22.3<br>3.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 26,562.5 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 15,082.34 ns</td>
      <td valign="middle">Valid<br>in 3,780.21 ns</td>
      <td valign="middle">Valid<br>in 56,282.98 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:::a2:a3:a4:b1:ffff:11.2<br>2.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,075.35 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 9,805.58 ns</td>
      <td valign="middle">Valid<br>in 3,366.95 ns</td>
      <td valign="middle">Valid<br>in 35,687.85 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:a3:a4::11.22.33.4<br>4]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,143.69 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 53,856.09 ns</td>
      <td valign="middle">Valid<br>in 2,299.32 ns</td>
      <td valign="middle">Valid<br>in 28,658.46 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2:a3:a4:b1::11.22.3<br>3.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,862.58 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 61,855.36 ns</td>
      <td valign="middle">Valid<br>in 1,885.59 ns</td>
      <td valign="middle">Valid<br>in 32,123.08 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::11.22.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,262.23 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 51,077.77 ns</td>
      <td valign="middle">Valid<br>in 1,579.25 ns</td>
      <td valign="middle">Valid<br>in 30,657.8 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1:a2::11.22.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,130.98 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 60,993.39 ns</td>
      <td valign="middle">Valid<br>in 1,778.34 ns</td>
      <td valign="middle">Valid<br>in 34,078.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:cdef::11<br>.22.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,571.26 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 62,354.45 ns</td>
      <td valign="middle">Valid<br>in 2,021.63 ns</td>
      <td valign="middle">Valid<br>in 28,234.01 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:0123:4567:89ab:CDEF::11<br>.22.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,373.72 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 54,878.62 ns</td>
      <td valign="middle">Valid<br>in 2,456.85 ns</td>
      <td valign="middle">Valid<br>in 24,725.04 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@[IPv6:a1::b2:11.22.33.44]<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,048.17 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 56,233.39 ns</td>
      <td valign="middle">Valid<br>in 1,663.08 ns</td>
      <td valign="middle">Valid<br>in 29,647.07 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">+@b.c<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,851.63 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 34,862.42 ns</td>
      <td valign="middle">Valid<br>in 629.16 ns</td>
      <td valign="middle">Valid<br>in 6,676.84 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">TEST@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,100.75 ns</td>
      <td valign="middle">Valid<br>in 7,774.66 ns</td>
      <td valign="middle">Valid<br>in 1,081.57 ns</td>
      <td valign="middle">Valid<br>in 18,127.13 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">1234567890@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,492.63 ns</td>
      <td valign="middle">Valid<br>in 16,191.42 ns</td>
      <td valign="middle">Valid<br>in 2,634.44 ns</td>
      <td valign="middle">Valid<br>in 22,646.44 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test-test@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,649.84 ns</td>
      <td valign="middle">Valid<br>in 13,244.83 ns</td>
      <td valign="middle">Valid<br>in 1,471.84 ns</td>
      <td valign="middle">Valid<br>in 24,789.26 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">t*est@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,019.96 ns</td>
      <td valign="middle">Valid<br>in 9,069.25 ns</td>
      <td valign="middle">Valid<br>in 1,270.76 ns</td>
      <td valign="middle">Valid<br>in 21,458.62 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">+1~1+@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,770.22 ns</td>
      <td valign="middle">Valid<br>in 8,769.31 ns</td>
      <td valign="middle">Valid<br>in 1,476.05 ns</td>
      <td valign="middle">Valid<br>in 20,084.57 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">{_test_}@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,439.17 ns</td>
      <td valign="middle">Valid<br>in 12,292.65 ns</td>
      <td valign="middle">Valid<br>in 1,727.58 ns</td>
      <td valign="middle">Valid<br>in 22,173.01 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@123.123.123.x123<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,272.59 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,384.93 ns</td>
      <td valign="middle">Valid<br>in 1,776.31 ns</td>
      <td valign="middle">Valid<br>in 19,901.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cdburgess+!#$%&amp;'*-/=?+_{}|~test@gmail.co<br>m<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 19,153.68 ns</td>
      <td valign="middle">Valid<br>in 21,640.01 ns</td>
      <td valign="middle">Valid<br>in 3,642.44 ns</td>
      <td valign="middle">Valid<br>in 55,785.23 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">valid@about.museum<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,545.66 ns</td>
      <td valign="middle">Valid<br>in 12,525.11 ns</td>
      <td valign="middle">Valid<br>in 1,704.8 ns</td>
      <td valign="middle">Valid<br>in 14,365.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a@bar<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,239.39 ns</td>
      <td valign="middle">Valid<br>in 7,079.98 ns</td>
      <td valign="middle">Valid<br>in 784 ns</td>
      <td valign="middle">Valid<br>in 12,003.14 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a-b@bar.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,917.88 ns</td>
      <td valign="middle">Valid<br>in 10,642.29 ns</td>
      <td valign="middle">Valid<br>in 1,188.58 ns</td>
      <td valign="middle">Valid<br>in 14,151.71 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">+@b.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 5,977.56 ns</td>
      <td valign="middle">Valid<br>in 10,034.68 ns</td>
      <td valign="middle">Valid<br>in 1,039.36 ns</td>
      <td valign="middle">Valid<br>in 10,294.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(foo\@bar)@iamcal.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,305.95 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 5,151.76 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 48,181.73 ns</td>
      <td valign="middle">Valid<br>in 49,795.61 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">(comment)test@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 15,089.38 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,840.36 ns</td>
      <td valign="middle">Valid<br>in 1,894.52 ns</td>
      <td valign="middle">Valid<br>in 61,531.89 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">(foo)cal(bar)@(baz)iamcal.com(quux)<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 16,655.85 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,032.71 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 47,338.56 ns</td>
      <td valign="middle">Valid<br>in 68,322.72 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(foo\)bar)@iamcal.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,400.67 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,056.31 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 38,152.97 ns</td>
      <td valign="middle">Valid<br>in 41,081.57 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal(woo(yay)hoopla)@iamcal.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,074.2 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,520.97 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 61,714.7 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 41,807.14 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(Welcome to
 the ("wonderf<br>ul" (!)) world
 of email)@test.<br>org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,542.87 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,486.3 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 48,214.52 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 83,510.63 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">pete(his account)@silly.test(his host)<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,468.43 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 2,085.34 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 36,383.08 ns</td>
      <td valign="middle">Valid<br>in 61,290.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(abc\(def)@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,996.43 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,734.5 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 32,763.45 ns</td>
      <td valign="middle">Valid<br>in 38,994.81 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a(a(b(c)d(e(f))g)h(i)j)@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,711.89 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,999.05 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 37,295.36 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 14,621.9 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">c@(Chris's host.)public.example<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,060.43 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,584.84 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 45,495.89 ns</td>
      <td valign="middle">Valid<br>in 23,954.66 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">a@b.co-foo.uk<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 12,205.44 ns</td>
      <td valign="middle">Valid<br>in 16,623.14 ns</td>
      <td valign="middle">Valid<br>in 3,139.33 ns</td>
      <td valign="middle">Valid<br>in 9,202.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">_Yosemite.Sam@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,924.18 ns</td>
      <td valign="middle">Valid<br>in 11,563.05 ns</td>
      <td valign="middle">Valid<br>in 1,774.61 ns</td>
      <td valign="middle">Valid<br>in 29,006.41 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">~@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,506.33 ns</td>
      <td valign="middle">Valid<br>in 8,956.7 ns</td>
      <td valign="middle">Valid<br>in 1,153.62 ns</td>
      <td valign="middle">Valid<br>in 8,415.67 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Iinsist@(that comments are allowed)this.<br>is.ok<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 18,821.4 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,035.3 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 42,324.11 ns</td>
      <td valign="middle">Valid<br>in 44,948.81 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@Bücher.ch<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 27,425.71 ns</td>
      <td valign="middle">Valid<br>in 26,080.78 ns</td>
      <td valign="middle">Valid<br>in 1,189.14 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 15,939.08 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">あいうえお@example.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,479.02 ns</td>
      <td valign="middle">Valid<br>in 13,133.34 ns</td>
      <td valign="middle">Valid<br>in 1,315.79 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,582.78 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Pelé@example.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 13,423.59 ns</td>
      <td valign="middle">Valid<br>in 12,083.89 ns</td>
      <td valign="middle">Valid<br>in 1,351.91 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 11,687.77 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">δοκιμή@παράδειγμα.δοκιμή<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 77,502.76 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 136,855.19 ns</td>
      <td valign="middle">Valid<br>in 9,409.59 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,582.45 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">我買@屋企.香港<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 20,399.77 ns</td>
      <td valign="middle">Valid<br>in 24,599.5 ns</td>
      <td valign="middle">Valid<br>in 4,932.75 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,924.19 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">二ノ宮@黒川.日本<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 23,751.54 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 64,151.68 ns</td>
      <td valign="middle">Valid<br>in 4,100.73 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 5,082.18 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">медведь@с-балалайкой.рф<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 128,287.85 ns</td>
      <td valign="middle">Valid<br>in 21,923.01 ns</td>
      <td valign="middle">Valid<br>in 7,012.56 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,081.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">संपर्क@डाटामेल.भारत<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 32,757 ns</td>
      <td valign="middle">Valid<br>in 28,814.02 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 29,493.29 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,271.47 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">email@example.com (Joe Smith)<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,372.87 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,342.02 ns</td>
      <td valign="middle">Valid<br>in 7,283.08 ns</td>
      <td valign="middle">Valid<br>in 25,398.03 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">cal@iamcal(woo).(yay)com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,694.71 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 7,874.91 ns</td>
      <td valign="middle">Valid<br>in 6,861.06 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 19,399.96 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(abc.def).last@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 7,057.18 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,693.78 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 36,778.52 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 43,793.58 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first(a"bc.def).last@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,803.16 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,379.09 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 29,224.71 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 38,610.52 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.(")middle.last(")@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,844.02 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,932.79 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 40,444.65 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 21,090.99 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first.last@x(123456789012345678901234567<br>8901234567890123456789012345678901234567<br>890).com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 19,380.84 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 19,050.64 ns</td>
      <td valign="middle">Valid<br>in 9,854.93 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 91,384.02 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">user%uucp!path@berkeley.edu<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,491.65 ns</td>
      <td valign="middle">Valid<br>in 14,527.34 ns</td>
      <td valign="middle">Valid<br>in 12,981.25 ns</td>
      <td valign="middle">Valid<br>in 23,040.91 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">first().last@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,749.52 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,462.13 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 38,022.49 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 27,123.99 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">mymail\@hello@hotmail.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,864.42 ns</td>
      <td valign="middle">Valid<br>in 13,264.96 ns</td>
      <td valign="middle">Valid<br>in 9,610.23 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 17,044.78 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Abc\@def@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,276.63 ns</td>
      <td valign="middle">Valid<br>in 7,848.77 ns</td>
      <td valign="middle">Valid<br>in 8,629.03 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 13,695.09 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Fred\ Bloggs@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,895.42 ns</td>
      <td valign="middle">Valid<br>in 13,325.75 ns</td>
      <td valign="middle">Valid<br>in 11,722.37 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 11,851.5 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe.\\Blow@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 14,002 ns</td>
      <td valign="middle">Valid<br>in 10,806.06 ns</td>
      <td valign="middle">Valid<br>in 9,495.76 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,009.64 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">{^c\@**Dog^}@cartoon.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,937.85 ns</td>
      <td valign="middle">Valid<br>in 13,285.97 ns</td>
      <td valign="middle">Valid<br>in 13,132.86 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 10,647.05 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">phil.h\@\@ck@haacked.com<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 17,082.38 ns</td>
      <td valign="middle">Valid<br>in 11,353.83 ns</td>
      <td valign="middle">Valid<br>in 12,080.42 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,806.93 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">abc\\@test.org<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 8,633.82 ns</td>
      <td valign="middle">Valid<br>in 5,926.58 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 45,629.78 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 12,563.17 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">Joe A Smith &lt;email@example.com&gt;<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,689.54 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 3,444.7 ns</td>
      <td valign="middle">Valid<br>in 9,502.44 ns</td>
      <td valign="middle">Valid<br>in 15,638.04 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">"John Michael" &lt;tester@test.net&gt;<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,374.18 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 4,808.55 ns</td>
      <td valign="middle">Valid<br>in 5,780.73 ns</td>
      <td valign="middle">Valid<br>in 13,244.81 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">A &lt;apple@pie.com&gt;<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 9,088.87 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,814.13 ns</td>
      <td valign="middle">Valid<br>in 5,283.23 ns</td>
      <td valign="middle">Valid<br>in 9,489.19 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">ABC.DEF@GHI.JKL (MNO)<br></th>
      <td valign="middle">Valid<br></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 11,029.65 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 1,142.13 ns</td>
      <td valign="middle">Valid<br>in 6,553.54 ns</td>
      <td valign="middle">Valid<br>in 25,093.06 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">xyz@🚀.kz<br></th>
      <td valign="middle">Valid<br><small class="text-muted">Emoji domains are valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 23,083.9 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 286,512.19 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 18,940.54 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 7,737.9 ns</td>
    </tr><tr>
      <th scope="row" valign="middle">test@😉.tld<br></th>
      <td valign="middle">Valid<br><small class="text-muted">Emoji domains are valid</small></td>
      <td style="background-color:#56666B" valign="middle">Valid<br>in 10,737.86 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 235,592.29 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 19,723.2 ns</td>
      <td style="background-color:#815355" valign="middle">Invalid<br>in 8,856.61 ns</td>
    </tr><tr>
      <th scope="row">Totals</th>
      <td>315/315</td>
      <td>315/315<br>Average Time: 10,864.62 ns</td>
      <td>245/315<br>Average Time: 21,434.04 ns</td>
      <td>218/315<br>Average Time: 17,713.09 ns</td>
      <td>216/315<br>Average Time: 26,995.25 ns</td>
  </tr>
"""

result = parse_html(html_string)

# Sort result map:
result = {k: v for k, v in sorted(
    result.items(), key=lambda item: (item[1], item[0]))}

text = ""
for key, value in result.items():
    # escape " for each key
    escaped_key = key.replace('"', '\\"')
    text += '\t"{}": {},\n'.format(escaped_key,
                                   "true" if value == True else "false")

# Save to file:
with open("result.txt", "w") as f:
    f.write(text)

print(text)
