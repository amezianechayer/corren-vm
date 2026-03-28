# aurex-vm

A virtual machine for moving money.

This repo bundles:
* The Aurex VM
* A FaRl parser
* A FaRl compiler

## Example

```farl
transfer [DZD.2 1099] from @users:001:wallet then @users:001:credit
send 85% to @drivers:033
send 15% to @platform:fees
```

## FaRl Syntax

**FaRl** stands for *Financial Accounting Rules Language*.

### Simple transfer

```farl
transfer [DZD.2 500] from @ameziane to @yanis
```

### Transfer with variables

```farl
var $buyer: account
var $seller: account
var $amount: monetary

transfer $amount from $buyer to $seller
```

### Send all balance

```farl
transfer [DZD.2 *] from @users:001 to @platform
```

### Other statements

```farl
print 100 + 50 - 25
fail
```

## Conventions

**Assets** — `ASSET.PRECISION` format

```
DZD.2   — Algerian Dinar, 2 decimal places
EUR.2   — Euro, 2 decimal places
SYM     — asset without precision
```

**Accounts** — always prefixed with `@`

```
@world          — unlimited source
@ameziane      — user account
@platform       — platform account
```

## Related

* [aurex](https://github.com/amezianechayer/aurex) — the programmable financial ledger