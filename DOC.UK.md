# is — довідник

Повний довідник пакета `is`: контракт валідації, усі функції за темами, бренди
`CardKind` і практичні рецепти.

Англійська версія: **[DOC.md](DOC.md)**.

## Зміст

- [Контракт валідації](#контракт-валідації)
- [Обмеження типів](#обмеження-типів)
- [Акаунт та ідентичність](#акаунт-та-ідентичність)
- [Фінанси](#фінанси)
- [Географія](#географія)
- [Мережа та зв'язок](#мережа-та-звязок)
- [Символьні класи рядків](#символьні-класи-рядків)
- [Числа](#числа)
- [Кодування та формат](#кодування-та-формат)
- [Мобільний зв'язок і телеком](#мобільний-звязок-і-телеком)
- [Помилки](#помилки)
- [Рецепти й поради](#рецепти-й-поради)

## Контракт валідації

Кожна функція `is` відповідає на одне питання «так/ні» про **формат** свого
вводу й повертає `bool` (кілька повертають додаткове значення чи помилку). Через
увесь пакет діють три правила:

1. **Жодного очищення, жодної нормалізації.** `is` перевіряє рівно те, що ви
   передали. Він ніколи не зрізає пробіли, не прибирає роздільники й не змінює
   регістр. Якщо ввід потребує очищення — очистіть його спершу (наприклад,
   `Weed`/`Preserve` з пакета `g`), а тоді валідуйте.
2. **Суворість — явна.** Там, де сенс мають і м'якше, і суворіше прочитання,
   суворіше вмикається окремо через кінцевий аргумент `strict ...bool`, тож
   поведінка за замовчуванням не дивує.
3. **Без хибнопозитивних за задумом.** v2 підтягнув чимало валідаторів
   (`Numeric`, `MD5`, `Base64`, `IPv4`, рядкові `Latitude`/`Longitude`,
   `IMEI`/`IMSI`), щоб звичайні слова й хибні значення більше не проходили.

```go
import "github.com/goloop/is/v2"
```

## Обмеження типів

В узагальнених валідаторах трапляються два обмеження:

| Обмеження | Дозволяє |
|-----------|----------|
| `Numerable`  | усі цілі й дробові типи |
| `Verifiable` | ціле, дробове, `string`, `rune` |

Деякі валідатори узагальнені над `string | float64` (географічні помічники) чи
`string | int64` (`IMEI`), тож одним викликом можна перевірити або розпарсене
число, або його текстову форму.

## Акаунт та ідентичність

```go
func Email(email string) bool
func Nickname(nickname string, strict ...bool) bool
func VariableName(v string, strict ...bool) bool          // псевдонім: Var
func VariableNameFor(v string, language string) (bool, error) // псевдонім: VarFor
func SelectorName(v string, strict ...bool) bool          // псевдонім: Sel
```

`Email` перевіряє формат адреси (та відхиляє поспіль `.`/`-` усередині мітки).
`Nickname` перевіряє ім'я користувача; у суворому режимі дозволений лише
консервативний набір символів. Він **не** зрізає пробіли навколо.

`VariableName` перевіряє загальний ідентифікатор; `VariableNameFor` валідує за
правилами конкретної мови **та її зарезервованими словами** (Go, Python, C#,
Dart, Bash, Elixir, Erlang, Julia, Objective-C, VB.NET, COBOL, Fortran, Prolog,
Eiffel, Assembly та інші), повертаючи [`ErrLanguageNotSupported`](#помилки) для
невідомої мови. `SelectorName` перевіряє ім'я селектора у стилі CSS. `Var`,
`VarFor` і `Sel` — короткі псевдоніми.

```go
is.Email("user@example.com")            // true
is.Nickname("user@123", true)           // false (суворо)
ok, err := is.VariableNameFor("class", "python") // false, nil — зарезервоване слово
```

## Фінанси

```go
func BankCard(str string, kinds ...CardKind) bool
func IBAN(iban string, strict ...bool) bool               // псевдонім: Iban
func IBANCountry(iban string, strict ...bool) (string, bool)
func CalculateIBANChecksum(iban string) int
```

`BankCard` перевіряє номер картки (довжина та Луна); передайте одну чи кілька
констант `CardKind`, щоб також вимагати конкретний бренд. Бренди — це непрозорі
константи, а не змінні регекси, тож валідацію не можна тихо зламати на весь
процес:

| Константи-бренди (вибірка) |
|----------------------------|
| `Visa`, `VisaElectron`, `MasterCard`, `Maestro`, `DiscoverCard`, `JCB`, `UnionPay`, `ChinaUnionPay`, `DinersClub` (+ `DinersClubCarteBlanche`, `DinersClubInternational`, `DinersClubUSAndCanada`), … |

`IBAN` перевіряє міжнародний номер рахунку (контрольна сума MOD-97); суворий
режим також вимагає точну для країни довжину. `IBANCountry` повертає дволітерний
код країни, коли IBAN валідний. `CalculateIBANChecksum` повертає значення MOD-97
як `int` (`-1` за недійсного символу).

```go
is.BankCard("4111111111111111")            // true
is.BankCard("4111111111111111", is.Visa)   // true
is.IBAN("DE89370400440532013000")          // true
code, ok := is.IBANCountry("DE89370400440532013000") // "DE", true
```

## Географія

```go
func Latitude[T string | float64](lat T) bool
func Longitude[T string | float64](lon T) bool
func Coordinates[T string | float64](lat, lon T) bool
```

Перевіряють широту (`-90..90`), довготу (`-180..180`) чи пару. Рядковий шлях
приймає лише звичайні десяткові — наукова нотація, шістнадцяткові float і
роздільники цифр відхиляються; шлях `float64` перевіряє числовий діапазон.

```go
is.Coordinates(51.5074, -0.1278) // true
is.Latitude("91.0")              // false
```

## Мережа та зв'язок

```go
func IPv4(ip string) bool     func IPv6(ip string) bool     func IP(ip string) bool
func MAC(v string) bool
func URL(v string) bool
func Hostname(v string) bool  func Domain(v string) bool
func Phone(phone string) bool func E164(v string) bool
```

`IPv4` збігається з dotted-quad адресами й (як `net/netip`) відхиляє провідні
нулі; `IPv6` — з IPv6; `IP` приймає будь-яку. `MAC` перевіряє адресу IEEE 802.
`URL` вимагає абсолютний URL (схема + хост). `Hostname` за RFC 1123; `Domain` —
це hostname з літерним TLD. `Phone` приймає поширені роздільники (пробіли,
дефіси, крапки, дужки); `E164` вимагає суворий `+` і до 15 цифр.

```go
is.IP("2001:db8::1")               // true
is.Domain("example.com")           // true
is.Phone("+1 (234) 567-8900")      // true
is.E164("+12345678900")            // true
```

## Символьні класи рядків

```go
func Alpha(s string) bool    func Alnum(s string) bool    func Digit(s string) bool
func Lower(s string) bool    func Upper(s string) bool    func Title(s string) bool
func Space(s string) bool
func Numeric(s string) bool  func Decimal(s string) bool  func Float(s string) bool
```

Предикати символьного класу над усім рядком. `Alpha` — лише літери, `Alnum` —
літери й цифри, `Digit` — ASCII-цифри, `Lower`/`Upper`/`Title` — регістр,
`Space` — пробіли. `Numeric` приймає лише справжні десяткові цифри (Unicode-
категорія `Nd`) — літерні нумерали, CJK/римські цифри й самотні знаки
відхиляються. `Decimal` і `Float` перевіряють числові літерали.

```go
is.Alnum("abc123")   // true
is.Numeric("Ⅻ")      // false (римська цифра, не Nd)
is.Float("3.14")     // true
```

## Числа

```go
func Even[T Numerable](v T, f ...bool) bool
func Odd[T Numerable](v T, f ...bool) bool
func Whole[T Numerable](v T) bool
func Natural[T Numerable](v T) bool
func Positive[T Numerable](v T) bool
func Negative[T Numerable](v T) bool
func Zero[T Numerable](v T) bool
```

Перевірки числових властивостей для будь-якого цілого чи дробового типу. `Whole`
повідомляє про цілісне значення, `Natural` — про додатне ціле, а
`Zero`/`Positive`/`Negative` — про знак.

```go
is.Even(4)       // true
is.Natural(5)    // true
is.Positive(-1)  // false
```

## Кодування та формат

```go
func Base64(v string) bool   func Base64URL(v string) bool
func Hex(v string) bool      func Bin(v string) bool
func HexColor(v string) bool func RGBColor(v string) bool
func MD5(v string) bool      func SHA1(v string) bool
func SHA256(v string) bool   func SHA512(v string) bool
func UUID(v string) bool     func JWT(v string) bool
```

Валідатори формату. `Base64` використовує стандартний алфавіт (`A–Z a–z 0–9 + /`);
`Base64URL` — URL-безпечний алфавіт. `Hex`/`Bin` перевіряють набір цифр.
`HexColor` (`#RRGGBB`) і `RGBColor` перевіряють кольори. Валідатори гешів
вимагають рівно потрібну кількість шістнадцяткових нібблів (MD5 = 32,
SHA-1 = 40, SHA-256 = 64, SHA-512 = 128) без префікса. `UUID` перевіряє
канонічну форму `8-4-4-4-12`; `JWT` — тришматкову форму токена.

```go
is.Base64("SGVsbG8=")   // true
is.HexColor("#FF5733")  // true
is.UUID("550e8400-e29b-41d4-a716-446655440000") // true
```

## Мобільний зв'язок і телеком

```go
func IMEI[T string | int64](imei T) bool
func IMSI(imsi string) bool
```

`IMEI` перевіряє 15-значний ідентифікатор обладнання (з перевіркою Луна) з рядка
чи `int64`; `IMSI` — ідентифікатор абонента. Обидва приймають лише цифри —
провідний знак відхиляється.

```go
is.IMEI("490154203237518") // true
is.IMSI("310150123456789") // true
```

## Помилки

```go
var ErrLanguageNotSupported = errors.New("programming language is not supported")
```

Повертається з `VariableNameFor`/`VarFor` для невідомої мови. Це сентинел, тож
звіряйте через `errors.Is`:

```go
if _, err := is.VariableNameFor(name, lang); errors.Is(err, is.ErrLanguageNotSupported) {
    // непідтримувана мова
}
```

## Рецепти й поради

**Очистити, тоді валідувати.** `is` ніколи не очищує ввід. Поєднайте його з `g`,
щоб спершу прибрати шум:

```go
raw := "GB82 WEST 1234 5698 7654 32"
iban := g.Weed(raw, g.Whitespaces) // прибрати пробіли
ok := is.IBAN(iban)                // валідувати
```

**Вмикайте суворий режим на межах довіри.** Для нікнеймів та ідентифікаторів від
користувача передавайте `true`, щоб приймався лише консервативний набір
символів; лишайте м'який дефолт для внутрішніх даних, яким уже довіряєте.

**Вимагайте бренд лише за потреби.** `BankCard(num)` приймає будь-який валідний
бренд; додавайте константи `CardKind`, лише коли ваш потік обмежений конкретними
мережами, — надмірне обмеження відхиляє інакше валідні картки.

**Перевіряйте значення, а не його текст, коли число вже є.** Узагальнені
`Latitude`/`Longitude`/`Coordinates` та `IMEI` приймають розпарсений тип
напряму, оминаючи перетворення в рядок і назад.
