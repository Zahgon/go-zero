package dartgen

const (
	varTemplate = `import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import '../data/tokens.dart';

/// store tokens to local
///
/// pass null will clean local stored tokens
/// returns true if success, otherwise false
Future<bool> setTokens(Tokens tokens) async {
  var sp = await SharedPreferences.getInstance();
  if (tokens == null) {
    sp.remove('tokens');
    return true;
  }
  return await sp.setString('tokens', jsonEncode(tokens.toJson()));
}

/// get local stored tokens
///
/// if no, returns null
Future<Tokens> getTokens() async {
  try {
    var sp = await SharedPreferences.getInstance();
    var str = sp.getString('tokens');
    if (str == null || str.isEmpty) {
      return null;
    }
    return Tokens.fromJson(jsonDecode(str));
  } catch (e) {
    print(e);
    return null;
  }
}
`

	varTemplateV2 = `import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import '../data/tokens.dart';

const String _tokenKey = 'tokens';

/// Saves tokens
Future<bool> setTokens(Tokens tokens) async {
  var sp = await SharedPreferences.getInstance();
  return await sp.setString(_tokenKey, jsonEncode(tokens.toJson()));
}

/// remove tokens
Future<bool> removeTokens() async {
  var sp = await SharedPreferences.getInstance();
  return sp.remove(_tokenKey);
}

/// Reads tokens
Future<Tokens?> getTokens() async {
  try {
    var sp = await SharedPreferences.getInstance();
    var str = sp.getString('tokens');
    if (str == null || str.isEmpty) {
      return null;
    }
    return Tokens.fromJson(jsonDecode(str));
  } catch (e) {
    print(e);
    return null;
  }
}`
)

func genVars(dir string, isLegacy bool, scheme string, hostname string) error {
	_ = "STUB: not implemented"
	return nil
}
