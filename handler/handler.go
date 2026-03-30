@@
-f, iactive := l.Next, l.Interactive()
+var queued []rune
+f, iactive := func() ([]rune, error) {
+	if queued != nil {
+		r := queued
+		queued = nil
+		return r, nil
+	}
+	return l.Next()
+}, l.Interactive()
 if iactive {
-	f = func() ([]rune, error) {
-		// next line
-		r, err := l.Next()
+	next := f
+	f = func() ([]rune, error) {
+		r, err := next()
 		if err != nil {
 			return nil, err
 		}
-		// save history
 		_ = l.Save(string(r))
+		if line, rest, ok := splitMetaCommandLine(r); ok {
+			queued = rest
+			return line, nil
+		}
 		return r, nil
 	}
+	} else {
+		next := f
+		f = func() ([]rune, error) {
+			r, err := next()
+			if err != nil {
+				return nil, err
+			}
+			if line, rest, ok := splitMetaCommandLine(r); ok {
+				queued = rest
+				return line, nil
+			}
+			return r, nil
+		}
 }
@@
 }
+
+func splitMetaCommandLine(r []rune) ([]rune, []rune, bool) {
+	start := 0
+	for start < len(r) && unicode.IsSpace(r[start]) {
+		start++
+	}
+	if start == len(r) || r[start] != '\\' {
+		return nil, nil, false
+	}
+	var quote rune
+	escaped := false
+	for i := start; i < len(r)-1; i++ {
+		c := r[i]
+		if quote != 0 {
+			if escaped {
+				escaped = false
+				continue
+			}
+			if c == '\\' {
+				escaped = true
+				continue
+			}
+			if c == quote {
+				quote = 0
+			}
+			continue
+		}
+		switch c {
+		case '\'', '"', '`':
+			quote = c
+		case '\\':
+			if r[i+1] != '\\' {
+				continue
+			}
+			if i == start || !unicode.IsSpace(r[i-1]) {
+				continue
+			}
+			if i+2 < len(r) && !unicode.IsSpace(r[i+2]) {
+				continue
+			}
+			line := []rune(strings.TrimRightFunc(string(r[:i]), unicode.IsSpace))
+			rest := []rune(strings.TrimLeftFunc(string(r[i+2:]), unicode.IsSpace))
+			return line, rest, true
+		}
+	}
+	return nil, nil, false
+}
