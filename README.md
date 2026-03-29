# corren-vm

A virtual machine for moving money.

This repo bundles:

- The Corren VM
- A FaRl parser (ANTLR4)
- A FaRl compiler

## Related

- [corren](https://github.com/amezianechayer/corren) — the programmable financial ledger

---

## FaRl — Financial accounting Rule Language

FaRl is the scripting language built into Corren. It expresses financial flows — transfers, splits, cascades — in a readable, type-safe syntax validated at compile time.

### Simple transfer

```farl
transfer [DZD.2 10000] (
  from @world
  to   @banque:reserve
)
```

### Revenue split

```farl
transfer [DZD.2 5000] (
  from @client:ameziane
  to   {
    90/100 to @vendeur:yanis
    10/100 to @plateforme:commission
  }
)
```

### Source cascade

```farl
# Drains wallet first, then credit, then falls back to @world
transfer [DZD.2 3000] (
  from {
    @wallet:ameziane
    @credit:ameziane
    @world
  }
  to @commande:0042
)
```

### Capped source

```farl
transfer [DZD.2 1000] (
  from max [DZD.2 500] from @alice
  to   @bob
)
```

### With remaining

```farl
transfer [DZD.2 5000] (
  from @client:ameziane
  to   {
    10%       to @plateforme:commission
    5%        to @taxes
    remaining to @vendeur:yanis
  }
)
```

### Transfer all balance

```farl
transfer [DZD.2 *] (
  from @escrow:commande:0042
  to   @vendeur:yanis
)
```

### With variables

```farl
var $montant:    monetary
var $vendeur:    account
var $commission: portion

transfer $montant (
  from @world
  to   {
    $commission to @plateforme
    remaining   to $vendeur
  }
)
```

### Metadata origin

```farl
var $sale:   account
var $seller: account = meta($sale,   "seller")
var $rate:   portion = meta($seller, "commission_rate")

transfer [DZD.2 10000] (
  from $sale
  to   {
    $rate     to @plateforme:commission
    remaining to $seller
  }
)
```

### Set metadata

```farl
set transaction metadata {
  "reference" = "PAY-2026-001"
  "channel"   = "mobile"
}

set account metadata of @vendeur:yanis key "commission_rate" = "12%"
```

### Other statements

```farl
print 100 + 50 - 25
fail
```

---

## Conventions

**Assets** — `ASSET` or `ASSET.PRECISION`

```
DZD.2   — Algerian Dinar, 2 decimal places
TND.3   — Tunisian Dinar, 3 decimal places
NGN.2   — Nigerian Naira, 2 decimal places
SYM     — asset without decimal precision
```

**Accounts** — always prefixed with `@`, namespaced with `:`

```
@world                   — unlimited source, no balance check
@ameziane                — simple account
@users:001:wallet        — namespaced account
@commande:1042:paiement  — nested namespace
```

**Variables** — prefixed with `$`

```
$montant    — monetary variable
$vendeur    — account variable
$commission — portion variable
```

**Portions**

```
85/100   — ratio form
15%      — percentage form
12.5%    — decimal percentage
remaining — the leftover after fixed portions
```

---

## Installation

```bash
git clone https://github.com/amezianechayer/corren-vm
cd corren-vm
go test ./...
```

**Requires:** Go 1.16+