# -*- mode: python ; coding: utf-8 -*-


a = Analysis(
    ['C:\\HORIZON_TEST\\Scripts\\horizon.py'],
    pathex=[],
    binaries=[],
    datas=[('C:\\HORIZON_TEST\\Xray-Core\\xray.exe', 'Xray-Core'), ('C:\\HORIZON_TEST\\Configs\\config.json', 'Configs'), ('C:\\HORIZON_TEST\\Logs', 'Logs')],
    hiddenimports=[],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    noarchive=False,
    optimize=0,
)
pyz = PYZ(a.pure)

exe = EXE(
    pyz,
    a.scripts,
    a.binaries,
    a.datas,
    [],
    name='horizon',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    upx_exclude=[],
    runtime_tmpdir=None,
    console=True,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)
