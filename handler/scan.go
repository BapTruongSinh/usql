@@
-		case '\\':
-			return scanMeta
+		case '\\':
+			if s.pos+1 < len(s.r) && s.r[s.pos+1] == '\\' && (s.pos == s.start || unicode.IsSpace(s.r[s.pos-1])) && (s.pos+2 == len(s.r) || unicode.IsSpace(s.r[s.pos+2])) {
+				s.pos += 2
+				return scanSQL
+			}
+			return scanMeta
