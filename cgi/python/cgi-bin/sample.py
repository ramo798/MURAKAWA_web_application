#少なくともcgiのimportは必須
import cgi
import cgitb
cgitb.enable()

print("Content-Type: text/html")
print()

#HTMLの値を取得
form = cgi.FieldStorage()

#forループでformに格納されている要素を全て出力
for key in form:
    value = form[key].value
    print('<p>%s: %s</p>' % (key, value))
