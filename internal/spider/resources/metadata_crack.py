# -*- coding: utf-8 -*-
# ---author:liang.du ---
# ---email:liangdu1992@gmail.com---
import time
import ctypes
import base64
import execjs
import random

def lsubid():
    haha = '''function lsu() {
		// const jsdom = require("jsdom");
		// const { JSDOM } = jsdom;
		var t = 402871197;
		function e(e) {
			e = typeof e === undefined || null === e ? '' : e['toString']();
			for (var r = 0; r < e['length']; r++) {
				var n = .02519603282416938 * (t += e['charCodeAt'](r));
				n -= t = n >>> 0,
				t = (n *= t) >>> 0,
				t += 4294967296 * (n -= t);
			}
			return 23283064365386964e-26 * (t >>> 0);
		}
		var r = e(' ') 
			, n = e(' ') 
			, i = e(' ') 
			, o = 1
			, a = [String('<div id="a-popover-root" style="z-index:-1;position:absolute;"></div>'), "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36", 
			new Date()['getTime']()];
			// new Date()['getTime']() 当前时间戳  1588070148637
		for (var u in a)
			a['hasOwnProperty'](u) && ((r -= e(a[u])) < 0 && (r += 1),
			(n -= e(a[u])) < 0 && (n += 1),
			(i -= e(a[u])) < 0 && (i += 1));
		function s(t) {
			return ('0000000000' + (4294967296 * (e = 2091639 * r + 23283064365386964e-26 * o,
			r = n,
			n = i,
			i = e - (o = 0 | e)))['toString']())['slice'](-t);
			var e;
		};
		return 'X' + s(2) + '-' + s(7) + '-' + s(7) + ':' + Math['floor'](new Date()['getTime']() / 1000);
	}'''
    js = execjs.compile(haha)
    lsu = js.call('lsu')
    return lsu

def Crack():
    starttime = int(time.time() * 1000)
    he = time.strftime("%Y-%m-%d-%H-%M-%S",time.localtime(starttime//1000))
    timeArray = time.strptime(he,"%Y-%m-%d-%H-%M-%S")
    ist = int(time.mktime(timeArray))
    lsu = lsubid()
    hahah = {'dupedPlugins': "Chrome PDF Plugin Chrome PDF Viewer Native Client ||1920-1080-1040-24-*-*-*",
             'errors': [],
             'flashVersion': 'null',
             'location': "https://www.amazon.com/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fsellercentral.amazon.com%2Fhome&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=sc_na_amazon_v2&openid.mode=checkid_setup&language=zh_CN&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&pageId=sc_na_amazon_v2&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&ssoResponse=eyJ6aXAiOiJERUYiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiQTI1NktXIn0.iYtE7Fv6aJ6v8chj2opIp_kOGZ33bdo5resZCKHOyv0Ovq0Egq5yMA.WdeSSQ_tHORwt2UX.rCFA9SmFW6sMqkq8A9L0t4rOE0SvBe993mELFt-VHyS636tvTYTD7NhHOrHZzB80D_qAjzY3KhcnpZEKTV7t_yZ-v0WIkpXgzr_GTOnAGCoq7uKI079hTMOVL-zZxFJswOXZSCQ7aC_uumC8RKta23jimSBYW9dJDKvfwnVJ7AKiKEjq2V6ZnOEUmPfPSTTYy_jbcPxHt4dmIEoc4g05St4Fat0ccd6kNcf6tb0YzM6zF8bwllfXv2Haslg7g9KT_oY2.3Vmm67W6l0WPcwF1ejudGw",
             'lsUbid': lsu,
             'metrics': {'tz': random.randint(0,9999), 'fp2': random.randint(10000,555555), 'lsubid': random.randint(0,9999), 'browser': random.randint(0,9999)},
             'plugins': "Chrome PDF Plugin Chrome PDF Viewer Native Client ||1920-1080-1040-24-*-*-*",
             'referrer': "",
             'screenInfo': "1920-1080-1040-24-*-*-*",
             'start': starttime,
             'timeZone': 8,
             'userAgent': "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36",
             'version': "4.0.0",
             'webDriver': 'null'}

    lha = ["{",
           str('"metrics":' + str(hahah['metrics'])) + ',',
           str('"start":' + str(hahah['start'])) + ',',
           str('"timeZone":' + str(hahah['timeZone'])) + ',',
           str('"lsUbid":"' + str(hahah['lsUbid'])) + '",',
           str('"referrer":"' + str(hahah['referrer'])) + '",',
           str('"userAgent":"' + str(hahah['userAgent'])) + '",',
           str('"location":"' + str(hahah['location'])) + '",',
           str('"webDriver":' + str(hahah['webDriver'])) + ',',
           str('"errors":' + str(hahah['errors'])) + ',',
           str('"version":"' + str(hahah['version'])) + '",',
           "}"]

    strhahah = ''
    for ha in lha:
        strhahah += ha

    print(strhahah)
    return strhahah

def crcTable():
    js = execjs.compile('''
		function crcTable(r){
			var crc = Array();
			for (var t=0; t<256; t++){
				for (var e=t, c=0; c<8;c++)
					1 == (1 & e) ? e = e >>> 1 ^ 3988292384 : e >>>= 1;
				crc[t] = e
			};
			return crc
		}''')
    crclist = js.call('crcTable')
    return crclist


def crc32_js(r):
    js = execjs.compile('''
		function crcTable(){
			var crc = Array();
			for (var t=0; t<256; t++){
				for (var e=t, c=0; c<8;c++)
					1 == (1 & e) ? e = e >>> 1 ^ 3988292384 : e >>>= 1;
				crc[t] = e;
			};
			return crc;
		};
		function crc(r){
		var t,e = 0;
		var crcl = crcTable();
		e ^= 4294967295;
		for (var c=0; c<String(r).length; c++)
			t = 255 & (e ^ String(r).charCodeAt(c));
			e = e >>> 8 ^ crcl[t];
		return 4294967295 ^ e};''')
    crclist = js.call('crcTable')
    # print(crclist)
    cr32 = js.call('crc',r)
    return cr32

def hex_js(t):
    js = execjs.compile('''function jst(t){return ['0123456789ABCDEF'.charAt(t >>> 28 & 15),'0123456789ABCDEF'.charAt(t >>> 24 & 15),'0123456789ABCDEF'.charAt(t >>> 20 & 15),'0123456789ABCDEF'.charAt(t >>> 16 & 15),'0123456789ABCDEF'.charAt(t >>> 12 & 15),'0123456789ABCDEF'.charAt(t >>> 8 & 15),'0123456789ABCDEF'.charAt(t >>> 4 & 15),'0123456789ABCDEF'.charAt(15 & t)]['join']('')};''')
    jst = js.call('jst',t)
    print(t,jst)
    return jst


def doEncrypt(r):
    js = execjs.compile('''function doE(r){
		t = [1888420705, 2576816180, 2347232058, 874813317];
		for (var e = Math['ceil'](r['length'] / 4), o = [], i = 0; i < e; i++)
			o[i] = (255 & r['charCodeAt'](4 * i)) + ((255 & r['charCodeAt'](4 * i + 1)) << 8) + ((255 & r['charCodeAt'](4 * i + 2)) << 16) + ((255 & r['charCodeAt'](4 * i + 3)) << 24);
		for (var n = Math['floor'](6 + 52 / e),a = o[0], c = o[e - 1], d = 0; n-- > 0; )
			for (var h = (d += 2654435769) >>> 2 & 3, u = 0; u < e; u++)
				a = o[(u + 1) % e],
				c = o[u] += (c >>> 5 ^ a << 2) + (a >>> 3 ^ c << 4) ^ (d ^ a) + (t[3 & u ^ h] ^ c);
		for (var f = [], s = 0; s < e; s++)
			f[s] = String['fromCharCode'](255 & o[s], o[s] >>> 8 & 255, o[s] >>> 16 & 255, o[s] >>> 24 & 255);
		return f['join']('');}''')
    uf = js.call('doE',r)
    return uf


def base64(uf):
    js = execjs.compile('''function en(e){
		var c = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
		e = String(e);
		for (var t, r, o, n, d = e.length % 3, h = "", i = -1, f = e.length - d; ++i < f; )
			t = e.charCodeAt(i) << 16,
			r = e.charCodeAt(++i) << 8,
			o = e.charCodeAt(++i),
			h += c.charAt((n = t + r + o) >> 18 & 63) + c.charAt(n >> 12 & 63) + c.charAt(n >> 6 & 63) + c.charAt(63 & n);
		return 2 == d ? (t = e.charCodeAt(i) << 8,
			r = e.charCodeAt(++i),
			h += c.charAt((n = t + r) >> 10) + c.charAt(n >> 4 & 63) + c.charAt(n << 2 & 63) + "=") : 1 == d && (n = e.charCodeAt(i),
			h += c.charAt(n >> 2) + c.charAt(n << 4 & 63) + "=="),
			h}''')
    bas = js.call('en',uf)
    print("ECdITeCs:"+bas)


if __name__ == "__main__":
    sthahah = Crack()
    crc32 = crc32_js(sthahah)
    hek = hex_js(crc32)
    nstrkey = hek + '#' + sthahah
    uf = doEncrypt(nstrkey)
    print("-ww->\n",uf)
    bas = base64(uf)
    print("--->\n",bas)

