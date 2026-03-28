# aurex-vm

The compiler and virtual machine for [FaRl](https://github.com/amezianechayer/aurex-vm), a domain-specific language for money movements.

## Overview

aurex-vm takes FaRl scripts, compiles them into bytecode and executes them to produce financial postings.

FaRl stands for **Financial Accounting Rules Language**.

## FaRl Syntax

**Simple transfer**
```
transfer [DZD.2 100] from @yanis to @ilyes
```

**Transfer with variables**
```
{
    var $rider: account
    var $driver: account
    var $value: monetary
}
transfer $value
    from $rider
    to $driver
```

**Other statements**
```
print 40 + 5 - 8
fail
```

## Conventions

**Assets** — `ASSET.PRECISION` format
```
DZD.2   — Algerian Dinar, 2 decimal places
EUR.2   — Euro, 2 decimal places
```

**Accounts** — always prefixed with `@`
```
@world      — unlimited source
@ameziane      — user account
@platform   — platform account
```

## Related

- [aurex](https://github.com/amezianechayer/aurex) — the programmable financial ledger
