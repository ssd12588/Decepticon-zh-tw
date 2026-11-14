#!/usr/bin/env python3
"""Apply the zh-TW literal map to the Ink CLI sources.

Only whole string literals / JSX text nodes / simple template literals that
appear verbatim as a key in MAP are rewritten. Identifiers, flags, paths,
status values and protected command examples are listed in PROTECTED and are
never touched, even if they accidentally match a key.

Line endings are preserved byte-for-byte: files are read with newline=""
and written back with newline="".

Usage:
    python tools/i18n/tui_zh.py            # dry run, prints counts
    python tools/i18n/tui_zh.py --apply    # write changes
"""

from __future__ import annotations

import argparse
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
SRC = ROOT / "clients" / "cli" / "src"

# Never translate these, even if a MAP key looks similar.
PROTECTED = {
    "/blue up 127.0.0.1:3000",
    "/export markdown ./findings.md",
    "/mode red approve",
    "/scope import <file> --confirm",
    "/unknown foo",
    "custom/<your CUSTOM_OPENAI_MODEL>",
    "lm_studio/<your LMSTUDIO_MODEL>",
    "ollama_chat/<your OLLAMA_CLOUD_MODEL>",
    "nmap -sV target",
    "running",
    "completed",
    "failed",
    "pending",
    "cancelled",
    "success",
    "error",
    "unknown",
    "enabled",
    "disabled",
    "standard",
    "plugins",
}

from tui_zh_map import MAP, PREFIX_RULES  # noqa: E402


def targets(include_tests: bool) -> list[Path]:
    return sorted(
        p
        for p in SRC.rglob("*")
        if p.suffix in {".ts", ".tsx"}
        and (include_tests or ".test." not in p.name)
    )


def build_rules() -> list[tuple[re.Pattern[str], str]]:
    rules: list[tuple[re.Pattern[str], str]] = []
    # Longest first so "Ctrl+O transcript" wins over a shorter overlapping key.
    for en in sorted(MAP, key=len, reverse=True):
        if en in PROTECTED:
            continue
        zh = MAP[en]
        quoted_double = re.compile(r'"' + re.escape(en) + r'"')
        quoted_single = re.compile(r"'" + re.escape(en) + r"'")
        jsx = re.compile(r">(" + re.escape(en) + r")<")
        tick = re.compile(r"`(" + re.escape(en) + r")`")
        rules.append((quoted_double, f'"{zh}"'))
        rules.append((quoted_single, f"'{zh}'"))
        rules.append((jsx, f">{zh}<"))
        rules.append((tick, f"`{zh}`"))
    for pattern, repl in PREFIX_RULES:
        rules.append((re.compile(pattern), repl))
    return rules


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--apply", action="store_true", help="write files (default: dry run)")
    parser.add_argument("--include-tests", action="store_true", help="also rewrite *.test.ts(x) assertions")
    args = parser.parse_args()

    rules = build_rules()
    total = 0
    touched = 0
    for path in targets(args.include_tests):
        original = path.read_text(encoding="utf-8", newline="")
        updated = original
        count = 0
        for pattern, repl in rules:
            updated, hits = pattern.subn(lambda _m, r=repl: r, updated)
            count += hits
        if updated != original:
            touched += 1
            total += count
            print(f"{path.relative_to(ROOT)}: {count}")
            if args.apply:
                with path.open("w", encoding="utf-8", newline="") as handle:
                    handle.write(updated)
    mode = "applied" if args.apply else "dry run"
    print(f"\n{mode}: {total} replacements in {touched} files")
    return 0


if __name__ == "__main__":
    sys.exit(main())
