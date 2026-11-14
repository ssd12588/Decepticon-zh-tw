#!/usr/bin/env python3
"""Rewrite English sub-string assertions in the Ink CLI tests to zh-TW.

Only files matching *.test.ts / *.test.tsx are touched. Each pair is applied
verbatim (plain str.replace) so the intent of every assertion stays visible in
this table — no regex, no guessing.

Usage:
    python tools/i18n/tests_zh.py            # dry run
    python tools/i18n/tests_zh.py --apply
"""

from __future__ import annotations

import argparse
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
SRC = ROOT / "clients" / "cli" / "src"

PAIRS: list[tuple[str, str]] = [
    # blue.ts
    ("HTTP record reached", "HTTP 紀錄送達接收端"),
    ("record reached the receiver", "紀錄送達接收端"),
    ("no record reached", "沒有紀錄送達接收端"),
    ("No active engagement", "沒有進行中的授權測試任務"),
    # convenience.ts (/new, /plan argument rejection)
    ('StringContaining "Usage:"', 'StringContaining "用法："'),
    # doctorReport.ts
    ("Engagement: not selected", "Engagement: 未選擇"),
    ("target none (unscoped)", "target none (未設定範圍)"),
    ("authorization not confirmed", "authorization 未確認"),
    # planMode.ts
    ("engagement workspace", "授權測試任務工作區"),
    ("no enforced machine block", "強制執行用的 machine 區塊"),
    ("not a valid scope document", "不是有效的範圍文件"),
    # scope.ts
    ("Refusing|already exists", "拒絕|已存在"),
    # roeValidation.ts
    ("does not match", "不一致"),
    # sessionDoctor.ts
    ("machine scope ready", "機器範圍就緒"),
    # footer hints
    ("follow-up to continue", "後續訊息以繼續"),
    ("/resume sessions", "/resume 工作階段"),
    ('"esc to dismiss"', '"esc：關閉"'),
    # convenience.ts uses expect.stringContaining(...)
    ('stringContaining("Usage:")', 'stringContaining("用法：")'),
    # doctorReport.ts combined labels
    (
        "execution default (no pin) · planning default (no pin)",
        "execution 預設（未鎖定） · planning 預設（未鎖定）",
    ),
    (
        "target 10.0.0.1 (direct_scope) · authorization confirmed",
        "target 10.0.0.1 (direct_scope) · authorization 已確認",
    ),
    # sessionDoctor.ts
    ('toContain("unavailable")', 'toContain("無法使用")'),
]


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--apply", action="store_true")
    args = parser.parse_args()

    total = 0
    for path in sorted(p for p in SRC.rglob("*.test.ts*") if p.is_file()):
        original = path.read_text(encoding="utf-8", newline="")
        updated = original
        hits = 0
        for old, new in PAIRS:
            if old in updated:
                hits += updated.count(old)
                updated = updated.replace(old, new)
        if updated != original:
            total += hits
            print(f"{path.relative_to(ROOT)}: {hits}")
            if args.apply:
                with path.open("w", encoding="utf-8", newline="") as handle:
                    handle.write(updated)
    print(f"\n{'applied' if args.apply else 'dry run'}: {total} assertions")
    return 0


if __name__ == "__main__":
    sys.exit(main())
