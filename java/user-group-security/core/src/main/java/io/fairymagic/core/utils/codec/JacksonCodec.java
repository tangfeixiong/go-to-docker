package io.fairymagic.core.utils.codec;

import java.io.IOException;

import com.fasterxml.jackson.core.JsonGenerationException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.fairymagic.core.ugs.domain.User;
import io.fairymagic.core.ugs.domain.jackson.CommonView;

public class JacksonCodec {
    final  static  ObjectMapper mapper = new ObjectMapper();

    public static String normoalizeUser(User user) throws JsonGenerationException, JsonMappingException, IOException {
//        try {
            String result = mapper.writerWithView(CommonView.Normal.class).writeValueAsString(user);
            return result;
//        } catch (JsonGenerationException e) {
//            e.printStackTrace();
//        } catch (JsonMappingException e) {
//            e.printStackTrace();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }
    }
}
